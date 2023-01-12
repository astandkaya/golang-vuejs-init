package models

import (
    "github.com/jinzhu/gorm"
)

type UserModel struct {
    gorm.Model
    UserName  string `form:"username" json:"username" binding:"required"`
    Password  string `form:"password" json:"password" binding:"required"`
}

type UserRepository interface {
    Create(user *UserModel)
    Find(id int) *UserModel
    All() *[]UserModel
    Exists(username string, password string) bool
}