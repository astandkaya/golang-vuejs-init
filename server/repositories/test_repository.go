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

func (r *TestRepo) Find(id int) *models.TestModel {
    result := &models.TestModel{}
    r.db.Find(result, id)

    return result
}

func (r *TestRepo) All() *[]models.TestModel {
    result := &[]models.TestModel{}
    r.db.Find(result)

    return result
}
