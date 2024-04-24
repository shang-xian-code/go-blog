package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
)

func GetPostPigeonhole() (*models.PigeonholeRes, error) {
	//查询所有的文章 进行月份的整理
	//查询所有的分类
	var posts []models.Post
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err = dao.GetPostAll()
	if err != nil {
		return nil, err
	}
	// 根据月份进行分类
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		createAt := post.CreateAt
		month := createAt.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	var pigeonholeRes = &models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        &pigeonholeMap,
	}
	return pigeonholeRes, nil
}
