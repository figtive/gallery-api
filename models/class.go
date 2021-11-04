package models

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID          string    `gorm:"primaryKey;type:varchar(10);"`
	Name        string    `gorm:"type:varchar(64);not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"-"`
}

type ClassOrmer interface {
	Insert(class Class) (id string, err error)
	GetOneByID(id string) (class Class, err error)
}

type classOrm struct {
	db *gorm.DB
}

func NewClassOrmer(db *gorm.DB) ClassOrmer {
	_ = db.AutoMigrate(&Class{})
	return &classOrm{db}
}

func (o *classOrm) Insert(class Class) (id string, err error) {
	result := o.db.Model(&Class{}).Create(&class)
	return class.ID, result.Error
}

func (o *classOrm) GetOneByID(id string) (class Class, err error) {
	result := o.db.Model(&Class{}).Where("id = ?", id).First(&class)
	return class, result.Error
}
