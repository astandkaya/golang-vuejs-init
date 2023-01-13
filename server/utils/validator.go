package utils

import (
    "github.com/go-playground/validator/v10"
    "github.com/jinzhu/gorm"

    "app/models"
)

var validate *validator.Validate
var db *gorm.DB

func ValidatorInit(db_arg *gorm.DB) {
    db = db_arg
}

func Validator() *validator.Validate {
    validate = validator.New()
    validate.RegisterValidation("uniq-user-email", uniqUserEmail)

    return validate
}

func uniqUserEmail(fl validator.FieldLevel) bool {
    cnt := 0
    model := &models.UserModel{
        Email: fl.Field().String(),
    }
    
    db.Where(model).Find(model).Count(&cnt)

    return cnt == 0
}
