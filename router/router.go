package router

import (
	"goblog/api"
	"goblog/views"
	"net/http"
)

func Router() {
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	http.HandleFunc("/", views.HTML.Index)
	// 分类页面 //http://localhost:8080/c/1  1参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	// 登录
	http.HandleFunc("/login", views.HTML.Login)
	// 具体文章 http://localhost:8080/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)
	// 写文章
	http.HandleFunc("/writing", views.HTML.Writing)
	// 归档界面
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
