package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

func DbInit() *gorm.DB {
	log.Info("Starting Database Connection")
	dbURI := "host=localhost user=admin password=password dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Panic("Failed to connect database with error " + err.Error())
	}
	return db
}
