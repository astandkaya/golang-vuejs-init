package models

import (
    "github.com/jinzhu/gorm"
)

type UserModel struct {
    gorm.Model
    Email  string `form:"email" json:"email" binding:"required"`
    Password  string `form:"password" json:"password" binding:"required"`
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
