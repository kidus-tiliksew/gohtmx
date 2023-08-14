package repository

import (
	"github.com/kidus-tiliksew/gohtmx/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (t *TodoRepository) Create(data *models.Todo) error {
	return t.DB.Create(data).Error
}

func (t *TodoRepository) Update(data *models.Todo) error {
	return t.DB.Updates(data).Error
}

func (t *TodoRepository) Get(ID uint) (*models.Todo, error) {
	var result *models.Todo
	if err := t.DB.First(&result, ID).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (t *TodoRepository) GetAll() ([]models.Todo, error) {
	var result []models.Todo
	
	if err := t.DB.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (t *TodoRepository) Delete(ID uint) error {
	return t.DB.Delete(&models.Todo{}, ID).Error
}
