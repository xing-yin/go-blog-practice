package router

import (
	"go-blog-practice/api"
	"go-blog-practice/views"
	"net/http"
)

func Router() {
	// 1.页面 views 2.api 数据(json) 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1.post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
