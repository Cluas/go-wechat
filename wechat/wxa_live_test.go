package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestWXAService_GetLiveInfo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/business/getliveinfo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
				"errcode": 0,
				"errmsg": "ok",
				"room_info": [
					{
						"name": "直播房间名",
						"roomid": 1,
						"cover_img": "http://mmbiz.qpic.cn/mmbiz_jpg/Rl1RuuhdstSfZa8EEljedAYcbtX3Ejpdl2et1tPAQ37bdicnxoVialDLCKKDcPBy8Iic0kCiaiaalXg3EbpNKoicrweQ/0?wx_fmt=jpeg",
						"live_status": 101,
						"start_time": 1568128900,
						"end_time": 1568131200,
						"anchor_name": "李四",
						"share_img": "http://mmbiz.qpic.cn/mmbiz_jpg/Rl1RuuhdstSfZa8EEljedAYcbtX3Ejpdlp0sf9YTorOzUbGF9Eib6ic54k9fX0xreAIt35HCeiakO04yCwymoKTjw/0?wx_fmt=jpeg",
						"goods": [
							{
								"cover_img": "http://mmbiz.qpic.cn/mmbiz_png/FVribAGdErI2PmyST9ZM0JLbNM48I7TH2FlrwYOlnYqGaej8qKubG1EvK0QIkkwqvicrYTzVtjKmSZSeY5ianc3mw/0?wx_fmt=png",
								"url": "pages/index/index.html",
								"price": 1100,
								"name": "fdgfgf"
							}
						]
					}
				],
				"total": 1
			}
		`)
	})
	got, _, err := client.WXA.GetLiveInfo(context.Background(), "token", &GetLiveInfoRequest{
		Start: 0,
		Limit: 10,
	})
	if err != nil {
		t.Errorf("WXA.GetLiveInfo retured err: %v", err)
	}
	want := &LiveInfo{
		RoomInfo: []*RoomInfo{
			{
				Name:       "直播房间名",
				RoomID:     1,
				CoverImg:   "http://mmbiz.qpic.cn/mmbiz_jpg/Rl1RuuhdstSfZa8EEljedAYcbtX3Ejpdl2et1tPAQ37bdicnxoVialDLCKKDcPBy8Iic0kCiaiaalXg3EbpNKoicrweQ/0?wx_fmt=jpeg",
				LiveStatus: 101,
				StartTime:  int64(1568128900),
				EndTime:    int64(1568131200),
				AnchorName: "李四",
				ShareImg:   "http://mmbiz.qpic.cn/mmbiz_jpg/Rl1RuuhdstSfZa8EEljedAYcbtX3Ejpdlp0sf9YTorOzUbGF9Eib6ic54k9fX0xreAIt35HCeiakO04yCwymoKTjw/0?wx_fmt=jpeg",
				Goods: []*Good{
					{
						CoverImg: "http://mmbiz.qpic.cn/mmbiz_png/FVribAGdErI2PmyST9ZM0JLbNM48I7TH2FlrwYOlnYqGaej8qKubG1EvK0QIkkwqvicrYTzVtjKmSZSeY5ianc3mw/0?wx_fmt=png",
						URL:      "pages/index/index.html",
						Price:    1100,
						Name:     "fdgfgf",
					},
				},
			},
		},
		Total: 1,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("WXA.GetLiveInfo got %+v, want %+v", got, want)
	}

}

func TestWXAService_GetLiveInfo_replay(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/business/getliveinfo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
					"live_replay": [
						{
							"expire_time": "2020-11-11T03:49:55Z",
							"create_time": "2019-11-12T03:49:55Z",
							"media_url": "http://xxxxx.vod2.myqcloud.com/xxxxx/xxxxx/f0.mp4"
						}
					],
					"errcode": 0,
					"total": 1,
					"errmsg": "ok"
				}
		`)
	})
	got, _, err := client.WXA.GetLiveInfo(context.Background(), "token", &GetLiveInfoRequest{
		Start: 0,
		Limit: 10,
	})
	if err != nil {
		t.Errorf("WXA.GetLiveInfo retured err: %v", err)
	}
	expireTime, _ := time.Parse(time.RFC3339, "2020-11-11T03:49:55Z")
	createTime, _ := time.Parse(time.RFC3339, "2019-11-12T03:49:55Z")
	want := &LiveInfo{
		LiveReplay: []*LiveReplay{
			{
				ExpireTime: expireTime,
				CreateTime: createTime,
				MediaURL:   "http://xxxxx.vod2.myqcloud.com/xxxxx/xxxxx/f0.mp4",
			},
		},
		Total: 1,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("WXA.GetLiveInfo got %+v, want %+v", got, want)
	}
}
