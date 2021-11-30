package dao

import (
	"github.com/satori/uuid"
	"helloworld/db"
	"helloworld/model"
)

func AddPost(post *model.Post) (*model.Post, error) {
	if err := db.DB.Create(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func GetPost() ([]*model.Post, error) {
	posts := []*model.Post{}
	if err := db.DB.Find(&posts).Error; nil != err {
		return nil, err
	}
	return posts, nil
}

func UpdatePost(post *model.Post) error {
	if err := db.DB.Save(&post).Error; nil != err {
		return err
	}
	return nil
}

func DeletePostById(id uuid.UUID) error {
	post := &model.Post{}
	if err := db.DB.Where("id=?", id).Delete(post).Error; nil != err {
		return err
	}
	return nil
}