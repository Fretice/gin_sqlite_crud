package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Student struct {
	gorm.Model
	Name     string
	Age      int
	BirthDay time.Time
	Gender   int
	Email    string
	Phone    string
	Address  string
}

func (s *Student) GetList(page int, pagesize int, db *gorm.DB) []Student {
	var students []Student
	db.Order("id desc").Offset((page - 1) * pagesize).Limit(pagesize).Find(&students)
	return students
}

func (s *Student) GetListTotalCount(db *gorm.DB) int {
	var students []Student
	var count int
	db.Find(&students).Count(&count)
	return count
}

func (s *Student) Search(searchkey string,page int, pagesize int, db *gorm.DB) []Student {
	var students []Student
	db.Where("name like ? ", "%"+searchkey+"%").Or("email like ?", "%"+searchkey+"%").Order("id desc").Offset((page - 1) * pagesize).Limit(pagesize).Find(&students)
	return students
}

func(s *Student) SearchResultCount(searchkey string, db *gorm.DB)int{
	var students []Student
	var count int
	db.Where("name like ? ", "%"+searchkey+"%").Or("email like ?", "%"+searchkey+"%").Find(&students).Count(&count)
	return count
}

func (s *Student) Create(model *Student, db *gorm.DB) {
	db.Create(&model)
}

func (s *Student) GetDetail(id int, db *gorm.DB) Student {
	var model Student
	db.Find(&model, id)
	return model
}

func (s *Student) Update(updatedModel *Student, db *gorm.DB) {
	db.Save(&updatedModel)
}

func (s *Student) Delete(id int, db *gorm.DB) {
	var model Student
	db.First(&model, id)
	db.Delete(&model)
}
