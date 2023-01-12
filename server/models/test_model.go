package models

import (
    "github.com/jinzhu/gorm"
)

type TestModel struct {
    gorm.Model
    Test    string
}

func Test() *TestModel {
    return &TestModel{
    }
}

type TestRepository interface {
    Find(id int) *TestModel
    All() *[]TestModel
}