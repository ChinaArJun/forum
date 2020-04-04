package main

import (
	"fmt"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()// 解析参数
	fmt.Println(r.Form)
	fmt.Println("URL:", r.URL.Path) // 在服务端打印请求参数
	fmt.Println("Scheme", r.URL.Scheme)

	fmt.Fprintf(w, "你好") // 发送响应给客户端
}

func main() {
	http.HandleFunc("/", sayHelloWorld)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ser :", err)
		return
	}
}