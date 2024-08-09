package database

import (
	"fossil/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// global connection pool
type ConnectionPool struct {
	DB     *gorm.DB
	models []interface{}
}

var Conn ConnectionPool

func Connect(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		return err
	}

	// set global db object to db connection pool
	Conn = ConnectionPool{
		DB: db,
		models: []interface{}{
			model.User{},
		},
	}

	return nil
}
