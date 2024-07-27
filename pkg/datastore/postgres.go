package datastore

import (
	"fmt"

	"github.com/harisapturr/go-echo-boilerplate/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(conf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai",
		conf.PostgresHost,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDBName,
		conf.PostgresPort,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
