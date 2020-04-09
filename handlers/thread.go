package handlers

import (
	"fmt"
	"github.com/forum/models"
	"github.com/forum/pkg/helper"
	"net/http"
)

// get /threads/new
// 创建群主页面
func NewThread(writer http.ResponseWriter, request *http.Request)  {
	_, err := helper.Session(writer, request)
	if err != nil {
		// 去登陆
		http.Redirect(writer, request, "/login", 302)
	} else {
		// 访问的页面
		helper.GenerateHTML(writer, nil, "layout", "auth.navbar","new.thread")
	}
}
// POST /thread/create
// 执行群主创建逻辑
func CreateThread(writer http.ResponseWriter, request *http.Request)  {
	sess, err := helper.Session(writer, request)
	if err != nil {
		// 未登录
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			fmt.Println("connot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("connot get user from session")
		}
		// 获取参数
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			fmt.Println("createthread err")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// get /thread/read
// 通过ID渲染指定群主页面
func ReadThread(writer http.ResponseWriter, request *http.Request)  {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		fmt.Println("connot read thread")
	} else {
		_, err := helper.Session(writer, request)
		// 获取用户
		if err != nil {
			// 渲染前端页面
			helper.GenerateHTML(writer, thread, "layout", "navbar", "thread")
		} else {
			helper.GenerateHTML(writer, thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}