package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./tmp.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	//对于get请求，参数都放在url上（query param），请求体中是没有数据的
	queryParam := r.URL.Query() //自动帮我们识别url中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) //在服务端打印客户端发送来的请求body
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/login/", f1)
	http.HandleFunc("/hello/", f2)
	// http.ListenAndServe("192.168.153.1:9090", nil)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
