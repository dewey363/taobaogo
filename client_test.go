package taobaogo



import (
	"fmt"
	"taobao-go/config"
	"testing"
)

func Test_Request(t *testing.T) {
	res, err := Request("GET", map[string]string{
		"method":  "taobao.trades.sold.get",
		"fields":  "tid,type,status,payment,orders,rx_audit_status",
		"session": config.SessionKey,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
