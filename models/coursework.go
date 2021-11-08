package models

import (
	"time"

	"gorm.io/gorm"
)

type Coursework struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CourseID  string    `gorm:"not null"`
	Course    Course    `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
}

type CourseworkOrmer interface {
	Insert(coursework Coursework) (id string, err error)
	GetOneByID(id string) (Coursework, error)
}

type courseworkOrm struct {
	db *gorm.DB
}

func NewCourseworkOrmer(db *gorm.DB) CourseworkOrmer {
	_ = db.AutoMigrate(&Coursework{})
	return &courseworkOrm{db}
}

func (o *courseworkOrm) Insert(coursework Coursework) (id string, err error) {
	result := o.db.Model(&Coursework{}).Create(&coursework)
	return coursework.ID, result.Error
}

func (o *courseworkOrm) GetOneByID(id string) (Coursework, error) {
	var coursework Coursework
	result := o.db.Model(&Coursework{}).Where("id = ?", id).First(&coursework)
	return coursework, result.Error
}
