package handler

import (
	"net/http"

	domains "github.com/harisapturr/go-echo-boilerplate/internal/customer/domain"
	"github.com/harisapturr/go-echo-boilerplate/internal/customer/model/dto"
	"github.com/harisapturr/go-echo-boilerplate/pkg/utils"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	usecase domains.CustomerUseCase
}

func NewCustomerHandler(e *echo.Echo, customerUseCase domains.CustomerUseCase) {
	handler := &CustomerHandler{usecase: customerUseCase}

	group := e.Group("customer")
	group.GET("/:email", utils.Wrap(handler.GetDetail))
	group.GET("", utils.Wrap(handler.GetList))
	group.POST("", utils.Wrap(handler.Create))
	group.PATCH("/:customer_id", utils.Wrap(handler.Update))
	group.DELETE("/:customer_id", utils.Wrap(handler.Delete))
}

func (h *CustomerHandler) GetDetail(c echo.Context) utils.Response {
	ctx := c.Request().Context()
	payload := new(dto.GetDetailRequest)

	if err := utils.BindAndValidate(c, payload); err != nil {
		return utils.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	result, err := h.usecase.GetDetail(ctx, *payload)
	if err != nil {
		return utils.Response{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return utils.Response{
		Data:       result,
		Message:    "Success get detail customer",
		StatusCode: http.StatusOK,
	}
}

func (h *CustomerHandler) GetList(c echo.Context) utils.Response {
	ctx := c.Request().Context()

	payload := dto.NewGetListRequest(c)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return utils.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	payload.CalculateOffset()
	result, totalData, err := h.usecase.GetList(ctx, *payload)
	if err != nil {
		return utils.Response{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return utils.Response{
		Data:       result,
		Message:    "Success get list customer",
		StatusCode: http.StatusOK,
		PaginationMeta: utils.PaginationMeta{
			Page:  payload.Page,
			Limit: payload.Limit,
			Total: totalData,
		},
	}
}

func (h *CustomerHandler) Create(c echo.Context) utils.Response {
	ctx := c.Request().Context()

	payload := new(dto.CreateRequest)
	if err := c.Bind(payload); err != nil {
		return utils.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	if err := c.Validate(payload); err != nil {
		return utils.Response{
			Message:    "Failed to validate request",
			Error:      err,
			StatusCode: http.StatusBadRequest,
		}
	}

	err := h.usecase.Create(ctx, *payload)
	if err != nil {
		return utils.Response{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return utils.Response{
		Data:       "",
		Message:    "Success create customer",
		StatusCode: http.StatusCreated,
	}
}

func (h *CustomerHandler) Update(c echo.Context) utils.Response {
	ctx := c.Request().Context()

	payload := new(dto.UpdateRequest)
	if err := c.Bind(payload); err != nil {
		return utils.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	if err := c.Validate(payload); err != nil {
		return utils.Response{
			Message:    "Failed to validate request",
			Error:      err,
			StatusCode: http.StatusBadRequest,
		}
	}

	err := h.usecase.Update(ctx, *payload)
	if err != nil {
		return utils.Response{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return utils.Response{
		Data:       "",
		Message:    "Success update customer",
		StatusCode: http.StatusOK,
	}
}

func (h *CustomerHandler) Delete(c echo.Context) utils.Response {
	ctx := c.Request().Context()

	payload := new(dto.DeleteRequest)
	if err := c.Bind(payload); err != nil {
		return utils.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	if err := c.Validate(payload); err != nil {
		return utils.Response{
			Message:    "Failed to validate request",
			Error:      err,
			StatusCode: http.StatusBadRequest,
		}
	}

	err := h.usecase.Delete(ctx, *payload)
	if err != nil {
		return utils.Response{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return utils.Response{
		Data:       "",
		Message:    "Success delete customer",
		StatusCode: http.StatusOK,
	}
}
