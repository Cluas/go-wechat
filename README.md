# go-wechat #

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Cluas/go-wechat)
[![Test Status](https://github.com/Cluas/go-wechat/workflows/tests/badge.svg)](https://github.com/Cluas/go-wechat/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/Cluas/go-wechat/branch/master/graph/badge.svg)](https://codecov.io/gh/Cluas/go-wechat)

微信开放平台小程序第三方接口支持

### 示例
```go
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

```
