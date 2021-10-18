package handlers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

type dbEntity struct {
	conn      *gorm.DB
	userOrmer models.UserOrmer
}

type module struct {
	db *dbEntity
}

type HandlerFunc interface {
	AuthParseGoogleJWT(jwtString string) (claims dtos.GoogleJWTClaim, err error)
	AuthGenerateJWT(userInfo dtos.User) (token string, err error)

	UserGetOneByEmail(email string) (userInfo dtos.User, err error)
	UserInsert(userInfo dtos.User) (id string, err error)
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
				conn:      db,
				userOrmer: models.NewUserOrmer(db),
			},
		}
		return
	}
}
