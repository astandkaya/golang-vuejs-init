package utils

import (
    "github.com/go-playground/validator/v10"
    "github.com/jinzhu/gorm"

    "app/models"
    "app/repositories"
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
    r := repositories.User(db)
    m := models.UserModel{
        Email: fl.Field().String(),
    }

    return !r.Exists(m)
}
