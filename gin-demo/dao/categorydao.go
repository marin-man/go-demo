package dao

import (
	"helloworld/db"
	"helloworld/model"
)

func AddCategory(category *model.Category) (*model.Category, error) {
	err := db.DB.Create(&category).Error
	if nil != err {
		return nil, err
	}
	return category, nil
}

func GetCategories() ([]*model.Category, error) {
	categories := []*model.Category{}
	err := db.DB.Find(&categories).Error
	if nil != err {
		return nil, err
	}
	return categories, err
}

func UpdateCategory(category *model.Category) error {
	err := db.DB.Save(&category).Error
	if nil != err {
		return err
	}
	return  nil
}

func DeleteCategoryById(id uint) error {
	category := &model.Category{}
	err := db.DB.Where("id=?", id).Delete(&category).Error
	if nil != err {
		return err
	}
	return nil
}