package datastore

import (
	"fmt"

	"github.com/harisapturr/go-echo-boilerplate/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(conf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		conf.MySQLUser,
		conf.MySQLPassword,
		conf.MySQLHost,
		conf.MySQLDBName,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
