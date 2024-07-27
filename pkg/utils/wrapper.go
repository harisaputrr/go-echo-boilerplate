package utils

import (
	"net/http"

	"github.com/harisapturr/go-echo-boilerplate/pkg/consts"
	"github.com/labstack/echo/v4"
)

const (
	StatusField   = "status"
	StatusSuccess = "success"
	StatusError   = "error"

	CodeField    = "code"
	DataField    = "data"
	MessageField = "message"
	ErrorField   = "errors"

	MetadataField  = "metadata"
	PageField      = "page"
	LimitField     = "limit"
	TotalField     = "total_data"
	TotalPageField = "total_page"
)

// Response body
type Response struct {
	StatusCode     int
	Message        string
	Error          error
	Data           interface{}
	PaginationMeta PaginationMeta
}

type PaginationMeta struct {
	Page  int
	Limit int
	Total int64
}

type EchoHandlerFunc func(c echo.Context) Response

func Wrap(fn EchoHandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := fn(c)
		return Translate(c, res)
	}
}

func Translate(c echo.Context, res Response) error {
	result := map[string]interface{}{
		StatusField:  StatusSuccess,
		MessageField: res.Message,
	}

	status := http.StatusOK
	if res.StatusCode > 0 {
		status = res.StatusCode
		result[CodeField] = status
	}

	if res.Error != nil {
		result[StatusField] = StatusError
		result[MessageField] = res.Error.Error()

		if errFields, ok := res.Error.(ValidationErrors); ok {
			result[ErrorField] = errFields.Errors
			result[MessageField] = consts.ErrBadRequest
		}
	}

	if res.Data != nil {
		result[DataField] = res.Data
	}

	if res.PaginationMeta != (PaginationMeta{}) {
		limit := res.PaginationMeta.Limit
		total := res.PaginationMeta.Total
		result[MetadataField] = map[string]interface{}{
			PageField:      res.PaginationMeta.Page,
			LimitField:     limit,
			TotalField:     total,
			TotalPageField: (total + int64(limit) - 1) / int64(limit),
		}
	}

	return c.JSON(status, result)
}
