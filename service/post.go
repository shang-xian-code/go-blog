package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
	"html/template"
	"log"
)

func GetPostDetail(pId int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pId)
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes, nil
}

func Writing() (*models.WritingRes, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var writingRes = &models.WritingRes{
		config.Cfg.Viewer.Title,
		config.Cfg.System.CdnURL,
		categorys,
	}
	return writingRes, nil
}

func SavePost(post *models.Post) error {
	err := dao.SavePost(post)
	return err
}

func UpdatePost(post *models.Post) error {
	err := dao.UpdatePost(post)
	return err
}

func SearchPost(condition string) ([]models.SearchResp, error) {
	posts, _ := dao.GetPostSearch(condition)
	var searchRes []models.SearchResp
	for _, post := range posts {
		searchRes = append(searchRes, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchRes, nil
}
