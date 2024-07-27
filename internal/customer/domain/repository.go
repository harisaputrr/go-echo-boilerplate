package domain

import (
	"context"

	"github.com/harisapturr/go-echo-boilerplate/internal/customer/model/dto"
	model "github.com/harisapturr/go-echo-boilerplate/internal/customer/model/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Insert(ctx context.Context, tx *gorm.DB, customer *model.Customer) (err error)
	FindByEmail(ctx context.Context, email string) (result *model.Customer, err error)
	FindByID(ctx context.Context, id int64) (result *model.Customer, err error)
	FindAll(ctx context.Context, payload dto.GetListRequest) (result []*dto.GetListResponse, err error)
	Count(ctx context.Context, payload dto.GetListRequest) (result int64, err error)
	Update(ctx context.Context, payload *model.Customer) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
