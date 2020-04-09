package handlers

import (
	"github.com/forum/models"
	. "github.com/forum/pkg/helper"
	"net/http"
)

// 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		GenerateHTML(w, threads, "layout", "navbar", "index")
		_, err := Session(w, r)
		if err != nil {
			GenerateHTML(w, threads, "layout", "navbar", "index")
		} else {
			GenerateHTML(w, threads, "layout", "auth.navbar", "index")
		}
	} else {
		GenerateHTML(w, threads, "layout", "navbar", "index")
	}

	//files := []string{"views/layout.html", "views/navbar.html", "views/index.html",}
	//templates := template.Must(template.ParseFiles(files...))
	//threads, err := models.Threads();
	//if err == nil {
	//	templates.ExecuteTemplate(w, "layout", threads)
	//}
}