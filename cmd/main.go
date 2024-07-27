package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/harisapturr/go-echo-boilerplate/pkg/datastore"
	"github.com/harisapturr/go-echo-boilerplate/pkg/middlewares"
	"github.com/harisapturr/go-echo-boilerplate/pkg/utils"

	"github.com/harisapturr/go-echo-boilerplate/config"

	customerHandler "github.com/harisapturr/go-echo-boilerplate/internal/customer/handler"
	customerRepository "github.com/harisapturr/go-echo-boilerplate/internal/customer/repository"
	customerUseCase "github.com/harisapturr/go-echo-boilerplate/internal/customer/usecase"

	// orderHandler "github.com/harisapturr/go-echo-boilerplate/internal/order/handlers"
	// orderRepository "github.com/harisapturr/go-echo-boilerplate/internal/order/repositories"
	// orderUseCase "github.com/harisapturr/go-echo-boilerplate/internal/order/usecases"

	// userRepository "github.com/harisapturr/go-echo-boilerplate/internal/user/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.LoadConfig()

	db := datastore.NewMySQL(conf)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           `[ROUTE] ${time_rfc3339} | ${status} | ${latency_human} ${latency} | ${method} | ${uri}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.Use(middlewares.RateLimit())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Validator = utils.NewCustomValidator()

	// setup repository
	customerRepository := customerRepository.NewCustomerRepository(db)
	// userRepository := userRepository.NewUserRepository(db)
	// OrderRepository := orderRepository.NewOrderRepository(db)

	// setup usecase
	customerUseCase := customerUseCase.NewCustomerUseCase(customerRepository, db)
	// orderUseCase := orderUseCase.NewOrderUseCase(OrderRepository, userRepository, db)

	// setup handler
	customerHandler.NewCustomerHandler(e, customerUseCase)
	// orderHandler.NewOrderHandler(e, orderUseCase)

	// setup default routes
	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello, 🌏! o(≧▽≦)o",
		})
	})

	// start server
	go func() {
		if err := e.Start(":" + conf.AppPort); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
