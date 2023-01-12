package repositories

import (
    "github.com/jinzhu/gorm"

    "app/models"
)

type UserRepository struct {
    db *gorm.DB
}

func User(db *gorm.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) Create(user *models.UserModel) {
    r.db.Create(user)
}

func (r *UserRepository) Find(id int) *models.UserModel {
    result := &models.UserModel{}
    r.db.Find(result, id)

    return result
}

func (r *UserRepository) All() *[]models.UserModel {
    result := &[]models.UserModel{}
    r.db.Find(result)

    return result
}

func (r *UserRepository) Exists(username string, password string) bool {
    result := &models.UserModel{}
    cnt := 0
    
    r.db.Where(&models.UserModel{
        UserName: username,
        Password: password,
    }).Find(result).Count(&cnt)

    return cnt > 0
}
