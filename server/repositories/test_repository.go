package repositories

import (
    "github.com/jinzhu/gorm"

    "app/models"
)

type TestRepository struct {
    db *gorm.DB
}

func Test(db *gorm.DB) *TestRepository {
    return &TestRepository{
        db: db,
    }
}

func (r *TestRepository) Find(id int) *models.TestModel {
    result := &models.TestModel{}
    r.db.Find(result, id)

    return result
}

func (r *TestRepository) All() *[]models.TestModel {
    result := &[]models.TestModel{}
    r.db.Find(result)

    return result
}
