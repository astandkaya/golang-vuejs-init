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