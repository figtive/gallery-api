package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Coursework struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CourseID       string    `gorm:"not null"`
	Course         Course    `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"-"`
	CourseworkType string    `gorm:"not null"`
}

type CourseworkOrmer interface {
	DeleteByID(id string) error
	GetOneByID(id string) (Coursework, error)
	GetManyByUserIDAndIsVotedJoinCourseworkType(userID, courseworkType string) ([]Coursework, error)
	Insert(coursework Coursework) (id string, err error)
	Update(coursework Coursework) error
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

func (o *courseworkOrm) GetManyByUserIDAndIsVotedJoinCourseworkType(userID, courseworkType string) ([]Coursework, error) {
	var courseworks []Coursework
	result := o.db.Model(&Coursework{}).Joins(fmt.Sprintf("inner join %[1]s on courseworks.id = %[1]s.coursework_id inner join votes on courseworks.id = votes.coursework_id", courseworkType)).Where("votes.user_id = ?", userID).Find(&courseworks)
	return courseworks, result.Error
}

func (o *courseworkOrm) Update(coursework Coursework) error {
	return o.db.Model(&Coursework{}).Omit("created_at").Where("id = ?", coursework.ID).Updates(coursework).Error
}

func (o *courseworkOrm) DeleteByID(id string) error {
	result := o.db.Model(&Coursework{}).Where("id = ?", id).Delete(&Coursework{})
	return result.Error
}
