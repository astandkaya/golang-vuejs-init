package repositories

import (
    "github.com/jinzhu/gorm"

    "app/models"
)

type TestRepo struct {
	db *gorm.DB
}

func Test(db *gorm.DB) *TestRepo {
	return &TestRepo{
		db: db,
	}
}

func (r *TestRepo) FindByID(ID int) (*models.TestModel, error) {
	return &models.TestModel{}, nil
}

func (r *TestRepo) Save(test *models.TestModel) error {
	return nil
}