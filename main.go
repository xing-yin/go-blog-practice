package main

import (
	"go-blog-practice/common"
	"go-blog-practice/router"
	"log"
	"net/http"
)

//type IndexData struct {
//	Title string `json:"title"`
//	Desc  string `json:"desc"`
//}

//// indexHtml 响应页面
//func indexHtml(w http.ResponseWriter, r *http.Request) {
//	t := template.New("index.html")
//	viewPath, _ := os.Getwd()
//	t, _ = t.ParseFiles(viewPath + "/template/index.html")
//	var indexData IndexData
//	indexData.Title = "ALanYin博客"
//	indexData.Desc = "这是入门教程"
//	err := t.Execute(w, indexData)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

func init() {
	// 模版加载
	common.LoadTemplate()
}

// 项目参考地址：https://mszlu.com/goblog/01.html#_1-%E6%90%AD%E5%BB%BA%E9%A1%B9%E7%9B%AE
func main() {
	// 程序的入口：一个项目只能有一个入口
	server := http.Server{
		Addr: "127.0.0.1:8002",
	}

	////http.HandleFunc("/index.html", indexHtml)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	// 路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
