package main

import (
	"context"
	"fmt"

	"github.com/Cluas/go-wechat/wechat"
)

func main() {
	client := wechat.NewClient(nil)
	testers, _, err := client.WXA.MemberAuth(context.Background(), "auth_token")
	if err != nil {
		// do something
		fmt.Println(testers)
	}
}
