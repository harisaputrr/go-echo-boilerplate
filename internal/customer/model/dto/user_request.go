package dto

import (
	"github.com/harisapturr/go-echo-boilerplate/pkg/utils"
	"github.com/labstack/echo/v4"
)

type (
	GetDetailRequest struct {
		Email string `param:"email" validate:"required,email"`
	}

	GetListRequest struct {
		Search string `json:"search" query:"search"`
		Sort   string `json:"sort"  query:"sort"`
		utils.DefaultPaginationAttributes
	}

	DeleteRequest struct {
		ID int64 `param:"customer_id"`
	}

	CreateRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Phone    string `json:"phone" validate:"required"`
		Address  string `json:"address" validate:"required"`
	}

	UpdateRequest struct {
		ID      int64  `param:"customer_id" validate:"required"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	ChangePasswordRequest struct {
		OldPassword string `json:"old_password" validate:"required"`
		NewPassword string `json:"new_password" validate:"required"`
	}
)

// Set default GetListRequest value
func NewGetListRequest(c echo.Context) *GetListRequest {
	return &GetListRequest{
		DefaultPaginationAttributes: utils.DefaultPaginationAttributes{
			Page:  1,
			Limit: 10,
		},
	}
}
