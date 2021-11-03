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
	conn            *gorm.DB
	blogOrmer       models.BlogOrmer
	classOrmer      models.ClassOrmer
	courseworkOrmer models.CourseworkOrmer
	projectOrmer    models.ProjectOrmer
	teamOrmer       models.TeamOrmer
	userOrmer       models.UserOrmer
}

type module struct {
	db *dbEntity
}

type HandlerFunc interface {
	AuthParseGoogleJWT(jwtString string) (claims dtos.GoogleJWTClaim, err error)
	AuthGenerateJWT(userInfo dtos.User) (token string, err error)

	BlogGetMany(skip int, limit int) (blogs []dtos.Blog, err error)
	BlogGetOne(id string) (blog dtos.Blog, err error)
	BlogInsert(blogInsert dtos.BlogInsert, classID string) (id string, err error)

	ClassGetOneByID(id string) (classInfo dtos.Class, err error)
	ClassInsert(classInfo dtos.Class) (id string, err error)

	CourseworkInsert(classID string) (id string, err error)

	ProjectGetOne(id string) (project dtos.Project, err error)
	ProjectGetMany(skip int, limit int) (projects []dtos.Project, err error)
	ProjectInsert(projectInfo dtos.ProjectInsert, classID string, thumbnailPath string) (id string, err error)

	TeamInsert(teamInfo dtos.TeamInsert) (id string, err error)

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
				conn:            db,
				blogOrmer:       models.NewBlogOrmer(db),
				classOrmer:      models.NewClassOrmer(db),
				courseworkOrmer: models.NewCourseworkOrmer(db),
				projectOrmer:    models.NewProjectOrmer(db),
				teamOrmer:       models.NewTeamOrmer(db),
				userOrmer:       models.NewUserOrmer(db),
			},
		}
		return
	}
}
