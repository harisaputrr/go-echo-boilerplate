package usecases

import (
	"context"

	"github.com/harisapturr/go-echo-boilerplate/internal/customer/domain"
	customerHelper "github.com/harisapturr/go-echo-boilerplate/internal/customer/helper"
	"github.com/harisapturr/go-echo-boilerplate/internal/customer/model/dto"
	model "github.com/harisapturr/go-echo-boilerplate/internal/customer/model/entity"

	// userdomain "github.com/harisapturr/go-echo-boilerplate/internal/user/domain"
	"github.com/harisapturr/go-echo-boilerplate/pkg/utils"

	"gorm.io/gorm"
)

type customerUseCase struct {
	repository domain.CustomerRepository
	// userRepository userdomain.UserRepository
	db *gorm.DB
}

func NewCustomerUseCase(repository domain.CustomerRepository, db *gorm.DB) domain.CustomerUseCase {
	return &customerUseCase{
		repository: repository,
		// userRepository: userRepository,
		db: db,
	}
}

func (u *customerUseCase) GetDetail(ctx context.Context, payload dto.GetDetailRequest) (result *model.Customer, err error) {
	result, err = u.repository.FindByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *customerUseCase) GetList(ctx context.Context, payload dto.GetListRequest) (result []*dto.GetListResponse, totalData int64, err error) {
	payload.Sort = customerHelper.GetListSortMapper(payload.Sort)

	resultCount, err := u.repository.Count(ctx, payload)
	if err != nil {
		return nil, 0, err
	}

	result, err = u.repository.FindAll(ctx, payload)
	if err != nil {
		return nil, 0, err
	}

	return result, resultCount, nil
}

func (u *customerUseCase) Create(ctx context.Context, payload dto.CreateRequest) (err error) {
	tx := u.db.Begin()

	customer := &model.Customer{
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.Phone,
		Address:     payload.Address,
	}

	err = u.repository.Insert(ctx, tx, customer)
	if err != nil {
		tx.Rollback()
		return err
	}

	// user := &userdomain.User{
	// 	Username:   payload.Username,
	// 	Password:   utils.HashPassword(payload.Password),
	// 	CustomerID: customer.ID,
	// }

	// err = u.userRepository.Insert(ctx, tx, user)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	tx.Commit()

	return nil
}

func (u *customerUseCase) Update(ctx context.Context, payload dto.UpdateRequest) (err error) {

	_, err = u.repository.FindByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	customer := &model.Customer{
		DefaultAttributes: utils.DefaultAttributes{
			ID: payload.ID,
		},
		Name:    payload.Name,
		Address: payload.Address,
	}
	err = u.repository.Update(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (u *customerUseCase) Delete(ctx context.Context, payload dto.DeleteRequest) (err error) {
	_, err = u.repository.FindByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	err = u.repository.Delete(ctx, payload.ID)
	if err != nil {
		return err
	}

	return nil
}
