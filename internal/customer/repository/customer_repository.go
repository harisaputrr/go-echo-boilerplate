package repositories

import (
	"context"
	"errors"

	"github.com/harisapturr/go-echo-boilerplate/internal/customer/domain"
	"github.com/harisapturr/go-echo-boilerplate/internal/customer/model/dto"
	model "github.com/harisapturr/go-echo-boilerplate/internal/customer/model/entity"
	"github.com/harisapturr/go-echo-boilerplate/pkg/utils"

	"gorm.io/gorm"
)

const (
	customerTableName = "customers"
)

type (
	customerRepository struct {
		db *gorm.DB
	}
)

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindByID(ctx context.Context, id int64) (result *model.Customer, err error) {
	err = r.db.Table(customerTableName).Where("id = (?)", id).First(&result).Error

	return
}

func (r *customerRepository) FindByEmail(ctx context.Context, email string) (result *model.Customer, err error) {
	err = r.db.Table(customerTableName).Where("email = (?)", email).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = utils.NewBadRequestError(errors.New("customer not found"))
	}

	return
}

func (r *customerRepository) FindAll(ctx context.Context, payload dto.GetListRequest) (result []*dto.GetListResponse, err error) {
	query := r.db.Table(customerTableName).Limit(payload.Limit).Offset(payload.Page)
	if payload.Search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+payload.Search+"%", "%"+payload.Search+"%")
	}
	if payload.Sort != "" {
		query = query.Order(payload.Sort)
	}
	err = query.Find(&result).Error

	return
}

func (r *customerRepository) Count(ctx context.Context, payload dto.GetListRequest) (result int64, err error) {
	query := r.db.Debug().Table(customerTableName)
	if payload.Search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+payload.Search+"%", "%"+payload.Search+"%")
	}
	err = query.Count(&result).Error

	return
}

func (r *customerRepository) Insert(ctx context.Context, tx *gorm.DB, customer *model.Customer) (err error) {
	err = tx.Table(customerTableName).Create(customer).Error

	return
}

func (r *customerRepository) Update(ctx context.Context, payload *model.Customer) (err error) {
	err = r.db.Table(customerTableName).Where("id = (?)", payload.ID).Updates(payload).Error

	return
}

func (r *customerRepository) Delete(ctx context.Context, id int64) (err error) {
	err = r.db.Table(customerTableName).Where("id = (?)", id).Delete(&model.Customer{}).Error

	return
}
