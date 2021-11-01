package models

import (
	"time"

	"gorm.io/gorm"
)

type Coursework struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
}

type CourseworkOrmer interface {
	Insert() (id string, err error)
}

type courseworkOrm struct {
	db *gorm.DB
}

func NewCourseworkOrmer(db *gorm.DB) CourseworkOrmer {
	_ = db.AutoMigrate(&Coursework{})
	return &courseworkOrm{db}
}

func (o *courseworkOrm) Insert() (id string, err error) {
	var coursework Coursework
	result := o.db.Model(&Coursework{}).Create(&coursework)
	return coursework.ID, result.Error
}
