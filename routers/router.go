package routers

import (
	"github.com/forum/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type WebRouter struct {
	Name 		string
	Method 		string
	Pattern 	string
	HandlerFunc http.HandlerFunc
}

// 声明webRouter切片存放所有web路由
type WebRouters []WebRouter

// 定义所有web路由
var webRouters = WebRouters{
	{	"home", "GEt", "/", handlers.Index},

}

func NewRouter() *mux.Router  {
	// 创建Mux.router路由
	router := mux.NewRouter().StrictSlash(true)
	// 遍历web.go中定义的所有 webRouters
	for _, route := range webRouters {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}
	return router
}