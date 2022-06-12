package views

import (
	"go-blog-practice/common"
	"go-blog-practice/config"
	"go-blog-practice/models"
	"net/http"
)

// index 假数据-显示首页内容
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	//页面上涉及到的所有的数据，必须有定义
	var categotys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categotys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	index.WriteData(w, hr)
}
