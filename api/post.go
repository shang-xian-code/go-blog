package api

import (
	"errors"
	"goblog/common"
	"goblog/dao"
	"goblog/models"
	"goblog/service"
	"goblog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	// 获取用户id
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录过期"))
	}
	uid := claim.Uid
	// POST save PUT update
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		categoryId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		articleType := float64(0)
		if params["type"] != nil {
			articleType = params["type"].(float64)
		}
		var post = &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       int(articleType),
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		if err := service.SavePost(post); err != nil {
			common.Error(w, errors.New("数据库错误"))
			return
		}
		common.Success(w, post)
		return
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		categoryId, _ := strconv.Atoi(params["categoryId"].(string))
		uId := params["userId"].(float64)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		articleType := float64(0)
		pid := params["pid"].(float64)
		if params["type"] != nil {
			articleType = params["type"].(float64)
		}
		var post = &models.Post{
			Pid:        int(pid),
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     int(uId),
			ViewCount:  0,
			Type:       int(articleType),
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		if err := service.UpdatePost(post); err != nil {
			common.Error(w, errors.New("数据库错误"))
			return
		}
		common.Success(w, post)
		return
	}
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/api/v1/post/")
	pId, _ := strconv.Atoi(id)
	post, err := dao.GetPostById(pId)
	if err != nil {
		common.Error(w, errors.New("数据库查询错误"))
	}
	common.Success(w, post)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp, err := service.SearchPost(condition)
	if err != nil {
		common.Error(w, errors.New("查询文章错误"))
		return
	}
	common.Success(w, searchResp)
}
