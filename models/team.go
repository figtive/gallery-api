package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(64);not null"`
	ClassID   string    `gorm:"column:class_id"`
	Class     Class     `gorm:"foreignKey:ClassID"`
	ProjectID string    `gorm:"column:project_id"`
	Project   Project   `gorm:"foreignKey:ProjectID"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

type TeamOrmer interface {
	Insert(team Team) (id string, err error)
}

type teamOrm struct {
	db *gorm.DB
}

func NewTeamOrmer(db *gorm.DB) TeamOrmer {
	_ = db.AutoMigrate(&Team{})
	return &teamOrm{db}
}

func (o *teamOrm) Insert(team Team) (id string, err error) {
	result := o.db.Model(&Team{}).Create(&team)
	return team.ID, result.Error
}
