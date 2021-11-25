package models

import (
	"time"

	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

type Blog struct {
	CourseworkID string     `gorm:"primaryKey"`
	Coursework   Coursework `gorm:"foreignKey:CourseworkID"`
	Title        string     `gorm:"not null;type:varchar(32)"`
	Author       string     `gorm:"not null"`
	Link         string     `gorm:"not null"`
	Category     string     `gorm:"column:category"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"-"`
}

type BlogOrmer interface {
	GetMany(skip int, limit int) (blogs []Blog, err error)
	GetManyByCourseIDAndTerm(courseID string, term, maxTerm time.Time) ([]Blog, error)
	GetManyByTermAndCourseIdSortByVotes(term time.Time, courseId string) ([]Blog, error)
	GetManyByUserIDJoinVote(userID string) ([]Blog, error)
	GetOneByCourseworkID(courseworkID string) (blog Blog, err error)
	GetManyBookmarkByUserID(userID string) ([]Blog, error)
	Insert(blog Blog) (id string, err error)
	Update(blog Blog) error
}

type blogOrm struct {
	db *gorm.DB
}

func NewBlogOrmer(db *gorm.DB) BlogOrmer {
	_ = db.AutoMigrate(&Blog{})
	return &blogOrm{db}
}

func (o *blogOrm) Insert(blog Blog) (id string, err error) {
	result := o.db.Model(&Blog{}).Create(&blog)
	return blog.CourseworkID, result.Error
}

func (o *blogOrm) GetOneByCourseworkID(courseworkID string) (blog Blog, err error) {
	result := o.db.Model(&Blog{}).Where("coursework_id = ?", courseworkID).Preload("Coursework").First(&blog)
	return blog, result.Error
}

// TODO: random ordering
func (o *blogOrm) GetMany(skip int, limit int) (blogs []Blog, err error) {
	result := o.db.Model(&Blog{}).Offset(skip).Preload("Coursework")
	if limit > 0 {
		result = result.Limit(limit)
	}
	result = result.Find(&blogs)
	return blogs, result.Error
}

func (o *blogOrm) GetManyByTermAndCourseIdSortByVotes(term time.Time, courseId string) ([]Blog, error) {
	var blogs []Blog
	result := o.db.
		Model(&Blog{}).
		Joins("INNER JOIN courseworks ON blogs.coursework_id = courseworks.id LEFT JOIN votes ON courseworks.id = votes.coursework_id").
		Where("blogs.created_at >= ? AND blogs.created_at < ? AND courseworks.course_id = ?", utils.TimeToTermTime(term), utils.NextTermTime(term), courseId).
		Order("Count(votes.id) DESC").
		Group("blogs.coursework_id").
		Preload("Coursework").
		Find(&blogs)
	return blogs, result.Error
}

func (o *blogOrm) GetManyBookmarkByUserID(userID string) ([]Blog, error) {
	var blogs []Blog
	result := o.db.
		Model(&Blog{}).
		Joins("INNER JOIN bookmarks ON blogs.coursework_id = bookmarks.coursework_id").
		Where("bookmarks.user_id >= ?", userID).
		Find(&blogs)
	return blogs, result.Error
}

func (o *blogOrm) GetManyByCourseIDAndTerm(courseID string, term, maxTerm time.Time) ([]Blog, error) {
	var blogs []Blog
	result := o.db.
		Model(&Blog{}).
		Joins("INNER JOIN courseworks ON blogs.coursework_id = courseworks.id").
		Where("courseworks.course_id = ? AND blogs.created_at >= ? AND blogs.created_at < ?", courseID, term, maxTerm).
		Preload("Coursework").
		Find(&blogs)
	return blogs, result.Error
}

func (o *blogOrm) GetManyByUserIDJoinVote(userID string) ([]Blog, error) {
	var blogs []Blog
	result := o.db.
		Model(&Blog{}).
		Joins("INNER JOIN votes ON blogs.coursework_id = votes.coursework_id").
		Where("votes.user_id = ?", userID).
		Preload("Coursework").
		Find(&blogs)
	return blogs, result.Error
}

func (o *blogOrm) Update(blog Blog) error {
	return o.db.Model(&Blog{}).Where("coursework_id = ?", blog.CourseworkID).Omit("created_at").Updates(blog).Error
}
