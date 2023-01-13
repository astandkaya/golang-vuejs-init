package models

import (
    "github.com/jinzhu/gorm"
)

type UserModel struct {
    gorm.Model
    Email  string `form:"email" json:"email" validate:"required,email,uniq-user-email"`
    Password  string `form:"password" json:"password" validate:"required"`
}

func User() *UserModel {
    return &UserModel{
    }
}

type UserRepository interface {
    Create(user *UserModel)
    Find(id int) *UserModel
    All() *[]UserModel
    Exists(user UserModel) bool
}
