package models

import (
    "github.com/jinzhu/gorm"
)

type TestModel struct {
    gorm.Model
    Test    string
}

type TestRepository interface {
    FindByID(ID int) (*TestModel, error)
    Save(test *TestModel) error
}