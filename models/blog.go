package models

import (
	"time"

	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

type Blog struct {
	CourseworkID string     `gorm:"primaryKey"`
	Coursework   Coursework `gorm:"foreignKey:CourseworkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title        string     `gorm:"not null;type:varchar(128)"`
	Author       string     `gorm:"not null"`
	Link         string     `gorm:"not null"`
	Category     string     `gorm:"column:category"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"-"`
}

type BlogOrmer interface {
	Insert(blog Blog) (string, error)
	GetMany(skip int, limit int, courseID, title, category string, startTerm, endTerm time.Time) ([]Blog, error)
	GetManyByTermAndCourseIDSortByVotes(term time.Time, courseID string) ([]Blog, error)
	GetManyByUserIDJoinVote(userID string) ([]Blog, error)
	GetOneByCourseworkID(courseworkID string) (blog Blog, err error)
	GetManyBookmarkByUserID(userID string) ([]Blog, error)
	Update(blog Blog) error
	DeleteByID(id string) error
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

func (o *blogOrm) GetMany(skip int, limit int, courseID, title, category string, startTerm, endTerm time.Time) (blogs []Blog, err error) {
	result := o.db.Model(&Blog{}).
		Where(Blog{Category: category}).
		Joins("INNER JOIN courseworks ON blogs.coursework_id = courseworks.id").
		Where("blogs.created_at >= ? AND blogs.created_at < ?", startTerm, endTerm)
	if courseID != "" {
		result = result.Where("courseworks.course_id = ?", courseID)
	}
	if title != "" {
		result = result.Where("LOWER(blogs.title) LIKE LOWER(?)", "%"+title+"%")
	}
	if category != "" {
		result = result.Where(Blog{Category: category})
	}
	if limit > 0 {
		result = result.Limit(limit)
	}
	result = result.Offset(skip).Preload("Coursework").Find(&blogs)
	return blogs, result.Error
}

func (o *blogOrm) GetManyByTermAndCourseIDSortByVotes(term time.Time, courseID string) ([]Blog, error) {
	var blogs []Blog
	result := o.db.
		Model(&Blog{}).
		Joins("INNER JOIN courseworks ON blogs.coursework_id = courseworks.id LEFT JOIN votes ON courseworks.id = votes.coursework_id").
		Where("blogs.created_at >= ? AND blogs.created_at < ? AND courseworks.course_id = ?", utils.TimeToTermTime(term), utils.NextTermTime(term), courseID).
		Group("blogs.coursework_id").
		Having("COUNT(votes.id) >= 1").
		Order("COUNT(votes.id) DESC").
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

func (o *blogOrm) GetManyBookmarkByUserID(userID string) ([]Blog, error) {
	var blogs []Blog
	result := o.db.
		Model(&Blog{}).
		Joins("INNER JOIN bookmarks ON blogs.coursework_id = bookmarks.coursework_id").
		Where("bookmarks.user_id = ?", userID).
		Find(&blogs)
	return blogs, result.Error
}

func (o *blogOrm) Update(blog Blog) error {
	return o.db.Model(&Blog{}).Where("coursework_id = ?", blog.CourseworkID).Omit("created_at").Updates(blog).Error
}

func (o *blogOrm) DeleteByID(id string) error {
	result := o.db.Model(&Blog{}).Where("coursework_id = ?", id).Delete(&Blog{})
	return result.Error
}
