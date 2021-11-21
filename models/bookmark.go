package models

import (
	"time"

	"gorm.io/gorm"
)

type Bookmark struct {
	ID           string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID       string     `gorm:"type:uuid;not null"`
	User         User       `gorm:"foreignkey:UserID"`
	CourseworkID string     `gorm:"type:uuid;not null"`
	Coursework   Coursework `gorm:"foreignkey:CourseworkID"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
}

type BookmarkOrmer interface {
	Insert(bookmark Bookmark) (string, error)
	GetOneByUserIDAndCourseworkID(userID, courseworkID string) (Bookmark, error)
	DeleteByUserIDAndCourseworkID(userID, courseworkID string) error
}

type bookmarkOrm struct {
	db *gorm.DB
}

func NewBookmarkOrmer(db *gorm.DB) BookmarkOrmer {
	_ = db.AutoMigrate(&Bookmark{})
	return &bookmarkOrm{db: db}
}

func (o *bookmarkOrm) Insert(bookmark Bookmark) (string, error) {
	result := o.db.Model(&Vote{}).Create(&bookmark)
	return bookmark.ID, result.Error
}

func (o *bookmarkOrm) GetOneByUserIDAndCourseworkID(userID, courseworkID string) (Bookmark, error) {
	var bookmark Bookmark
	result := o.db.Model(&Bookmark{}).Where("user_id = ? AND coursework_id = ?", userID, courseworkID).First(&bookmark)
	return bookmark, result.Error
}

func (o *bookmarkOrm) DeleteByUserIDAndCourseworkID(userID, courseworkID string) error {
	return o.db.Model(&Bookmark{}).Where("user_id = ? AND coursework_id = ?", userID, courseworkID).Delete(&Bookmark{}).Error
}
