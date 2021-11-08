package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	CourseworkID string     `gorm:"primaryKey"`
	Coursework   Coursework `gorm:"foreignKey:CourseworkID"`
	Author       string     `gorm:"not null"`
	Title        string     `gorm:"not null;type:varchar(32)"`
	Link         string     `gorm:"not null"`
	Category     string     `gorm:"column:category"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"-"`
}

type BlogOrmer interface {
	Insert(blog Blog) (id string, err error)
	GetOneByCourseworkID(courseworkID string) (blog Blog, err error)
	GetMany(skip int, limit int) (blogs []Blog, err error)
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
