package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	CourseworkID string     `gorm:"primaryKey"`
	Coursework   Coursework `gorm:"foreignKey:CourseworkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name         string     `gorm:"type:varchar(32);not null"`
	Team         string     `gorm:"not null"`
	Description  string     `gorm:"not null"`
	Thumbnail    string     `gorm:"not null"`
	Field        string     `gorm:"type:varchar(32);not null"`
	Active       bool       `gorm:"not null"`
	Metadata     string     `gorm:"not null"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"-"`
}

type ProjectOrmer interface {
	GetOneByCourseworkID(courseworkID string) (project Project, err error)
	GetMany(skip int, limit int) (projects []Project, err error)
	Insert(project Project) (courseworkID string, err error)
	UpdateThumbnail(courseworkID string, path string) (err error)
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
	result := o.db.Model(&Project{}).Where("coursework_id = ?", courseworkID).Preload("Coursework").First(&project)
	return project, result.Error
}

// TODO: random ordering
func (o *projectOrm) GetMany(skip int, limit int) (projects []Project, err error) {
	result := o.db.Model(&Project{}).Offset(skip).Preload("Coursework")
	if limit > 0 {
		result = result.Limit(limit)
	}
	result = result.Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) UpdateThumbnail(courseworkID string, path string) (err error) {
	result := o.db.Model(&Project{}).Where("coursework_id = ?", courseworkID).Update("thumbnail", path)
	return result.Error
}
