package utils

import (
	"time"

	"gorm.io/gorm"
)

type DefaultAttributes struct {
	ID        int64           `json:"id" gorm:"primary_key"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	CreatedAt *time.Time      `json:"created_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type DefaultPaginationAttributes struct {
	Page              int  `json:"page" query:"page"`
	Limit             int  `json:"limit" query:"limit" validate:"min=1"`
	DisablePagination bool `json:"disable_pagination" query:"disable_pagination"`
}

func (d *DefaultPaginationAttributes) CalculateOffset() {
	if d.DisablePagination {
		d.Page = 0
	}
	d.Page = d.Limit * (d.Page - 1)
}
