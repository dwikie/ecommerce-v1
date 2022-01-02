package services

import (
	"ecommerce-v1/auth/configs"
	"ecommerce-v1/auth/entity"
	"ecommerce-v1/auth/models"
	"ecommerce-v1/auth/utils"
)

func SignIn(model *models.SignIn) (*entity.Account, error) {
	var account *entity.Account
	var err error

	err = configs.DB.Where("email = ?", model.Email).First(&account).Error
	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(model.Password, account.Password)
	if err != nil {
		return nil, err
	}

	return account, nil
}
