package dto

import (
	"github.com/harisapturr/go-echo-boilerplate/pkg/utils"
)

type (
	GetListResponse struct {
		utils.DefaultAttributes
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
	}
)
