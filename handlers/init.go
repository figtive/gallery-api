package handlers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
)

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn *gorm.DB
}

type HandlerFunc interface {
}

var Handler HandlerFunc

func InitializeHandler() (err error) {
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		configs.AppConfig.DbHost, configs.AppConfig.DbPort, configs.AppConfig.DbName, configs.AppConfig.DbUser, configs.AppConfig.DbPassword,
	)), &gorm.Config{})

	if err != nil {
		log.Println("[INIT] failed connecting to PostgreSQL")
		return
	} else {
		log.Println("[INIT] connected to PostgreSQL")
		Handler = &module{
			db: &dbEntity{
				conn: db,
			},
		}
		return
	}
}
