package domain

import (
	"context"

	"github.com/harisapturr/go-echo-boilerplate/internal/customer/model/dto"
	model "github.com/harisapturr/go-echo-boilerplate/internal/customer/model/entity"
)

type CustomerUseCase interface {
	Create(ctx context.Context, payload dto.CreateRequest) (err error)
	GetDetail(ctx context.Context, payload dto.GetDetailRequest) (result *model.Customer, err error)
	GetList(ctx context.Context, payload dto.GetListRequest) (result []*dto.GetListResponse, totalData int64, err error)
	Update(ctx context.Context, payload dto.UpdateRequest) (err error)
	Delete(ctx context.Context, payload dto.DeleteRequest) (err error)
}
