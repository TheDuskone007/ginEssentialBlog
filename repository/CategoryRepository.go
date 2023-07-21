package repository

import (
	"ginEssential2/common"
	"ginEssential2/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

func (cr CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{Name: name}
	if err := cr.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := cr.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr CategoryRepository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := cr.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr CategoryRepository) DeleteById(id int) error {
	if err := cr.DB.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
