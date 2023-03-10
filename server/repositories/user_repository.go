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
    result := models.User()
    r.db.Find(result, id)

    return result
}

func (r *UserRepository) All() *[]models.UserModel {
    result := &[]models.UserModel{}
    r.db.Find(result)

    return result
}

func (r *UserRepository) Exists(user models.UserModel) bool {
    cnt := 0
    
    r.db.Where(user).Find(models.User()).Count(&cnt)

    return cnt > 0
}

func (r *UserRepository) Delete(user models.UserModel) {
    r.db.Where(user).Delete(user)
}
