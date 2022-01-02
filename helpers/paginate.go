package pagination

import (
	"ecommerce-v1/auth/models"
	"math"

	"gorm.io/gorm"
)

func Source(db *gorm.DB) *paginateConfig {
	cfg := new(paginateConfig)
	cfg.db = db
	return cfg
}

type paginateConfig struct {
	db    *gorm.DB
	model interface{}
	limit int64
	page  int64
	order string
	where string
}

func (p *paginateConfig) Model(model interface{}) *paginateConfig {
	p.model = model
	return p
}

func (p *paginateConfig) OrderBy(query string) *paginateConfig {
	p.order = query
	return p
}

func (p *paginateConfig) Where(query string) *paginateConfig {
	p.where = query
	return p
}

func (p *paginateConfig) Page(page int64) *paginateConfig {
	p.page = page
	return p
}

func (p *paginateConfig) Limit(limit int64) *paginateConfig {
	p.limit = limit
	return p
}

func (p *paginateConfig) Exec() (*models.HttpResponsePagination, error) {
	var model models.HttpResponsePagination
	var err error
	// Set default value page
	if p.page == 0 {
		p.page = 1
	}
	// Set default value limit
	if p.limit == 0 {
		p.limit = 10

	}

	err = p.db.Where(p.where).Model(p.model).Count(&model.TotalData).Error
	if err != nil {
		return nil, err
	}

	offset := (p.page - 1) * p.limit
	err = p.db.Offset(int(offset)).Limit(int(p.limit)).Where(p.where).Order(p.order).Find(&p.model).Error
	if err != nil {
		return nil, err
	}

	model.Page = p.page
	model.Limit = p.limit
	model.Success = true
	model.Error = nil
	model.Data = &p.model
	model.TotalPages = int64(math.Ceil(float64(model.TotalData) / float64(model.Limit)))

	return &model, nil
}
