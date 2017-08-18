package taobaogo

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

var (
	AppKey    string = ""
	AppSecret string = ""
	Router    string = ""
)

//Request 发送API调用请求
func Request(method string, params map[string]string) (res *simplejson.Json, err error) {
	err = checkConfig()
	if err != nil {
		return
	}
	params["method"] = method
	req, err := http.NewRequest("POST", Router, strings.NewReader(mkPostData(params)))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpclient := &http.Client{}
	httpclient.Timeout = time.Second * 3

	response, err := httpclient.Do(req)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("request error, [code:%d] [status:%s] %x", response.StatusCode, response.Status, response.Body)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	res, err = simplejson.NewJson(body)
	if err != nil {
		return
	}
	if errorResponse, ok := res.CheckGet("error_response"); ok {
		errbts, _ := errorResponse.Encode()
		err = errors.New("Request ERROR,INFO:" + string(errbts))
	}
	return
}

func sign(args url.Values) string {
	keys := []string{}
	for k := range args {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	pstr := ""
	for _, k := range keys {
		pstr += k + args.Get(k)
	}
	sign := md5.Sum([]byte(AppSecret + pstr + AppSecret))
	return strings.ToUpper(hex.EncodeToString(sign[:]))
}

func checkConfig() error {
	if AppKey == "" {
		return errors.New("AppKey未配置")
	}
	if AppSecret == "" {
		return errors.New("AppSecret未配置")
	}
	if Router == "" {
		return errors.New("Router未配置")
	}
	return nil
}

func defaultArgs() url.Values {
	args := url.Values{}
	//默认参数
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	args.Add("timestamp", timestamp)
	args.Add("format", "json")
	args.Add("app_key", AppKey)
	args.Add("v", "2.0")
	args.Add("sign_method", "md5")
	return args
}

func mkPostData(params map[string]string) string {
	args := defaultArgs()
	for key, val := range params {
		args.Set(key, val)
	}
	args.Add("sign", sign(args))
	return args.Encode()
}
