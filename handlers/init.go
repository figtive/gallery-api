package handlers

import (
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

type dbEntity struct {
	conn            *gorm.DB
	blogOrmer       models.BlogOrmer
	bookmarkOrmer   models.BookmarkOrmer
	courseOrmer     models.CourseOrmer
	courseworkOrmer models.CourseworkOrmer
	projectOrmer    models.ProjectOrmer
	userOrmer       models.UserOrmer
	voteOrmer       models.VoteOrmer
}

type module struct {
	db *dbEntity
}

type HandlerFunc interface {
	AuthParseGoogleJWT(jwtString string) (claims dtos.GoogleJWTClaim, err error)
	AuthGenerateJWT(userInfo dtos.User) (token string, err error)

	BlogDelete(id string) error
	BlogGetMany(skip int, limit int) (blogs []dtos.Blog, err error)
	BlogGetManyByCourseIDInCurrentTerm(courseID string, currentOnly bool) ([]dtos.Blog, error)
	BlogGetOne(id string) (blog dtos.Blog, err error)
	BlogInsert(blogInsert dtos.BlogInsert, courseID string) (id string, err error)
	BlogUpdate(blogInfo dtos.BlogUpdate) error

	BookmarkInsert(bookmark dtos.Bookmark) (string, error)
	BookmarkHasMarked(bookmark dtos.Bookmark) (bool, error)
	BookmarkDelete(bookmark dtos.Bookmark) error
	BookmarkGetManyBlogByUserID(userID string) ([]dtos.Blog, error)
	BookmarkGetManyProjectByUserID(userID string) ([]dtos.Project, error)

	CourseDelete(id string) error
	CourseGetAll() ([]dtos.Course, error)
	CourseGetOneByID(id string) (courseInfo dtos.Course, err error)
	CourseInsert(courseInfo dtos.Course) (id string, err error)
	CourseUpdate(courseInfo dtos.CourseUpdate) error

	CourseworkGetOneByID(id string) (dtos.Coursework, error)
	CourseworkGetVoted(userID, cwTyoe string) ([]dtos.Coursework, error)
	CourseworkInsert(courseID, courseworkType string) (string, error)

	LeaderboardBlog(term time.Time, courseID string) ([]dtos.Blog, error)
	LeaderboardProject(term time.Time, courseID string) ([]dtos.Project, error)

	ProjectDelete(id string) error
	ProjectDeleteThumbnail(id string, thumbnailPath string) error
	ProjectGetOne(id string) (project dtos.Project, err error)
	ProjectGetMany(skip int, limit int, name, field string) (projects []dtos.Project, err error)
	ProjectGetManyByCourseID(courseID string, currentOnly bool) ([]dtos.Project, error)
	ProjectInsert(projectInfo dtos.ProjectInsert, courseID string) (id string, err error)
	ProjectInsertThumbnail(id string, header *multipart.FileHeader) error
	ProjectUpdate(projectInfo dtos.ProjectUpdate) error

	UserGetOneByEmail(email string) (userInfo dtos.User, err error)
	UserInsert(userInfo dtos.User) (id string, err error)
	UserUpdate(userInfo dtos.User) (err error)

	VoteCountByCourseworkID(courseworkID string) (int64, error)
	VoteCountByUserIDJoinCourseworkType(userID, courseworkType string) (int64, error)
	VoteCountVoteByUserForCourseworkTypeInCourse(userID, courseID, courseworkType string, term time.Time) (int64, error)
	VoteHasVoted(userID, courseworkID string) (bool, error)
	VoteGetVotedBlogs(userID string) ([]dtos.Blog, error)
	VoteGetVotedProjects(userID string) ([]dtos.Project, error)
	VoteInsert(userID string, voteInfo dtos.VoteInsert) (string, error)
	VoteUnvote(userID, courseworkID string) error
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
				bookmarkOrmer:   models.NewBookmarkOrmer(db),
				courseOrmer:     models.NewCourseOrmer(db),
				courseworkOrmer: models.NewCourseworkOrmer(db),
				projectOrmer:    models.NewProjectOrmer(db),
				userOrmer:       models.NewUserOrmer(db),
				voteOrmer:       models.NewVoteOrmer(db),
			},
		}
		return
	}
}
