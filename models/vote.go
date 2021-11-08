package models

import (
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
	Insert(vote Vote) (string, error)
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
