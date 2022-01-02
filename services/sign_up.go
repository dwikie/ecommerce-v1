package services

import (
	"database/sql"
	"ecommerce-v1/auth/configs"
	"ecommerce-v1/auth/entity"
	"ecommerce-v1/auth/models"
	"ecommerce-v1/auth/utils"
	"time"

	"github.com/google/uuid"
)

func SignUp(model *models.SignUp) (*models.Account, error) {
	hashPassword := utils.HashPassword(model.Password)

	account := entity.Account{
		Id:        uuid.New().String(),
		Fullname:  model.Fullname,
		Birthdate: model.Birthdate,
		Email:     model.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: &sql.NullTime{},
	}

	err := configs.DB.Create(&account).Error

	if err != nil {
		return nil, err
	}

	data := &models.Account{
		Id:        account.Id,
		Fullname:  account.Fullname,
		Birthdate: account.Birthdate,
		Email:     account.Email,
		Phone:     account.Phone,
		CreatedAt: account.CreatedAt,
	}

	return data, nil
}
