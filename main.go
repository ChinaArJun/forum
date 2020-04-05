package main

import (
	// . 别名，通过这种方式引入的包可以直接调用包中对外可见的变量、方法和结构体，而不需要加上包名前缀。
	. "github.com/forum/routers"
	"log"
	"net/http"
)

func main() {
	startWebServer("9090")
}

func startWebServer(port string)  {
	r := NewRouter()

	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)  // 通过router统一管理
	log.Println("start http server" + port)
	err := http.ListenAndServe(":" + port, nil) // 启动协程监听请求，每个请求会创建一个goruntime监控

	if err != nil {
		log.Println("err listenAndServe :", err)
	}
}