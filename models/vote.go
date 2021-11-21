package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Vote struct {
	ID           string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID       string     `gorm:"type:uuid;not null"`
	User         User       `gorm:"foreignkey:UserID"`
	CourseworkID string     `gorm:"type:uuid;not null"`
	Coursework   Coursework `gorm:"foreignkey:CourseworkID"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
}

type VoteOrmer interface {
	CountByCourseworkID(courseworkID string) (int64, error)
	CountByUserIDJoinCourseworkType(userID, courseworkType string) (int64, error)
	Insert(vote Vote) (string, error)
	GetManyByUserIDCourseworkIDAndCreatedAt(userID, courseworkID string, createdAt time.Time) ([]Vote, error)
	GetOneByUserIDAndCourseworkID(userID, courseworkID string) (Vote, error)
	DeleteByUserIDAndCourseworkID(userID, courseworkID string) error
}

type voteOrm struct {
	db *gorm.DB
}

func NewVoteOrmer(db *gorm.DB) VoteOrmer {
	_ = db.AutoMigrate(&Vote{})
	return &voteOrm{db: db}
}

func (o *voteOrm) Insert(vote Vote) (string, error) {
	result := o.db.Model(&Vote{}).Create(&vote)
	return vote.ID, result.Error
}

func (o *voteOrm) GetManyByUserIDCourseworkIDAndCreatedAt(userID, courseworkID string, createdAt time.Time) ([]Vote, error) {
	var votes []Vote
	result := o.db.Model(&Vote{}).Where("user_id = ? AND coursework_id = ? AND created_at > ?", userID, courseworkID, createdAt).Find(&votes)
	return votes, result.Error
}

func (o *voteOrm) CountByCourseworkID(courseworkID string) (int64, error) {
	var count int64
	result := o.db.Model(&Vote{}).Where("coursework_id = ?", courseworkID).Count(&count)
	return count, result.Error
}

func (o *voteOrm) CountByUserIDJoinCourseworkType(userID, courseworkType string) (int64, error) {
	var count int64
	result := o.db.Model(&Vote{}).Joins(fmt.Sprintf("inner join %[1]s on votes.coursework_id = %[1]s.coursework_id", courseworkType)).Where("votes.user_id = ?", userID).Count(&count)
	return count, result.Error
}

func (o *voteOrm) GetOneByUserIDAndCourseworkID(userID, courseworkID string) (Vote, error) {
	var vote Vote
	result := o.db.Model(&Vote{}).Where("user_id = ? AND coursework_id = ?", userID, courseworkID).First(&vote)
	return vote, result.Error
}

func (o *voteOrm) DeleteByUserIDAndCourseworkID(userID, courseworkID string) error {
	return o.db.Model(&Vote{}).Where("user_id = ? AND coursework_id = ?", userID, courseworkID).Delete(&Vote{}).Error
}
