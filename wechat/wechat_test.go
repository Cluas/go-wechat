package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, tearDown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	client = NewClient(nil)
	client.BaseURL, _ = url.Parse(server.URL)
	serverURL = server.URL
	tearDown = server.Close
	return
}
func testURLParseError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error to be return")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}
func TestCheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(`{"errcode": -1,"errmsg": "system error"}`)),
	}
	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Errorf("Expected error response.")
	}

	want := &ErrorResponse{
		Response: res,
		Message:  "system error",
		Code:     -1,
	}
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

func TestDo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	client.Do(context.Background(), req, body)

	want := &foo{A: "a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}
func TestDo_nilContext(t *testing.T) {
	client, _, _, tearDown := setup()
	defer tearDown()

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(nil, req, nil)

	if !reflect.DeepEqual(err, errors.New("context must be non-nil")) {
		t.Errorf("Expected context must be non-nil error")
	}
}

func TestDo_httpError(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	resp, err := client.Do(context.Background(), req, nil)

	if err != nil {
		t.Fatal("Expected HTTP 400 error, got no error.")
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected HTTP 400 error, got %d status code.", resp.StatusCode)
	}
}

func TestErrorResponse_Error(t *testing.T) {
	res := &http.Response{Request: &http.Request{}}
	err := ErrorResponse{
		Response: res,
		Message:  "message",
		Code:     0,
	}
	if err.Error() == "" {
		t.Errorf("Expected non-empty ErrorResponse.Error()")
	}
}
func TestNewClint(t *testing.T) {
	c := NewClient(nil)
	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClint BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClinet UserAgent is %v, want %v", got, want)
	}

	cNew := NewClient(nil)
	if c.client == cNew.client {
		t.Error("NewClient returned same http.Client, but they should differ.")
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &BasicInfo{AppID: String("a")}, `{"app_id":"a"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient(nil)

	type T struct {
		A map[interface{}]interface{}
	}

	_, err := c.NewRequest("GET", ".", &T{})
	if err == nil {
		t.Errorf("Expected error to be returned.")
	}
	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error; got %#v.", err)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}
