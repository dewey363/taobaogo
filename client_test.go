package taobaogo

import (
	"fmt"
	"testing"

	"github.com/luobosoft/taobao-go/config"
)

func Test_Request(t *testing.T) {
	AppKey = "test"
	AppSecret = "test"
	Router = "https://gw.api.tbsandbox.com/router/rest"
	res, err := Request("taobao.trades.sold.get", map[string]string{
		"fields":  "tid,type,status,payment,orders,rx_audit_status",
		"session": config.SessionKey,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
