package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDatabase(cfg Configuration) *gorm.DB {
	// TODO: supported with logger
	db, err := gorm.Open(mysql.Open(cfg.MySqlDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
	})

	if err != nil {
		log.Fatalf("Error creating database connection")
	} else {
		log.Println("Database is connected")
	}

	return db
}
