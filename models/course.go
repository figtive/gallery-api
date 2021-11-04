package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          string    `gorm:"primaryKey;type:varchar(10);"`
	Name        string    `gorm:"type:varchar(64);not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"-"`
}

type CourseOrmer interface {
	Insert(course Course) (id string, err error)
	GetOneByID(id string) (course Course, err error)
}

type courseOrm struct {
	db *gorm.DB
}

func NewCourseOrmer(db *gorm.DB) CourseOrmer {
	_ = db.AutoMigrate(&Course{})
	return &courseOrm{db}
}

func (o *courseOrm) Insert(course Course) (id string, err error) {
	result := o.db.Model(&Course{}).Create(&course)
	return course.ID, result.Error
}

func (o *courseOrm) GetOneByID(id string) (course Course, err error) {
	result := o.db.Model(&Course{}).Where("id = ?", id).First(&course)
	return course, result.Error
}
