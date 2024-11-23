package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init_db_conn() {

	dsn := "host=localhost user=admin dbname=festtolist port=5432 sslmode=disable"
	println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
	}
	db.AutoMigrate(&User{}, &Query{})
}
