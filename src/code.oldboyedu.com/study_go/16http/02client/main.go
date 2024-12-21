package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 公用一个client适用于 请求比较频繁
// var (
// 	client = http.Client{
// 		Transport: &http.Transport{
// 			DisableKeepAlives: false,
// 		},
// 	}
// )

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/hello/?name=周琳&age=18")
	// if err!=nil{
	// 	fmt.Println("get url failed,err:",err)
	// 	return
	// }

	data := url.Values{} //url values
	urlObj, _ := url.Parse("http://127.0.0.1:9090/hello/")
	data.Set("name", "周琳")
	data.Set("age", "9000")
	queryStr := data.Encode() //url encode之后的url
	fmt.Println(queryStr)

	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	fmt.Println(urlObj.String())

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("get url failed,err:", err)
	// 	return
	// }

	//请求不是特别频繁，用完就关闭该链接
	//禁用KeepAlive的client
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body) //在客户端读出服务端返回的响应的body
	if err != nil {
		fmt.Printf("read resp. body failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
