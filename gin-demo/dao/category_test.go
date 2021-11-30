package dao

import (
	"helloworld/model"
	"testing"
	"time"
)

func TestAddCategory(t *testing.T) {
	c := &model.Category{
		Name: "Mathematics",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	AddCategory(c)
}

func TestUpdateCategory(t *testing.T) {

}