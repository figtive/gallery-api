package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	CourseworkID string     `gorm:"primaryKey"`
	Coursework   Coursework `gorm:"foreignKey:CourseworkID"`
	Name         string     `gorm:"type:varchar(32);not null"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"-"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"-"`
}

type ProjectOrmer interface {
	Insert(project Project) (courseworkID string, err error)
	GetOneByCourseworkID(courseworkID string) (project Project, err error)
}

type projectOrm struct {
	db *gorm.DB
}

func NewProjectOrmer(db *gorm.DB) ProjectOrmer {
	_ = db.AutoMigrate(&Project{})
	return &projectOrm{db}
}

func (o *projectOrm) Insert(project Project) (courseworkID string, err error) {
	result := o.db.Model(&Project{}).Create(&project)
	return project.CourseworkID, result.Error
}

func (o *projectOrm) GetOneByCourseworkID(courseworkID string) (project Project, err error) {
	result := o.db.Model(&Project{}).Where("coursework_id = ?", courseworkID).First(&project)
	return project, result.Error
}
