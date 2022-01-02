package repository

import (
	"ecommerce-v1/auth/configs"
	paginate "ecommerce-v1/auth/helpers"
	"ecommerce-v1/auth/models"
)

type account struct{}

var Account *account

func (acc *account) GetAllAccount(page, limit int64, where, order string) (*models.HttpResponsePagination, error) {
	var model []*models.Account
	data, err := paginate.Source(configs.DB).Model(model).Limit(limit).Page(page).Where(where).OrderBy(order).Exec()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (acc *account) GetAccountByEmail(email string) (*models.Account, error) {
	var model models.Account
	err := configs.DB.Where("email = ?", email).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (acc *account) GetAccountById() {}

func (acc *account) GetAllAccountLike() {}

func (acc *account) UpdatePhoneNumber() {}

func (acc *account) UpdatePassword() {}

func (acc *account) UpdateEmail() {}

func (acc *account) UpdateBirthdate() {}

func (acc *account) UpdateFullname() {}
