package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

type Project struct {
	CourseworkID string         `gorm:"primaryKey"`
	Coursework   Coursework     `gorm:"foreignKey:CourseworkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name         string         `gorm:"type:varchar(128);not null"`
	Team         string         `gorm:"not null"`
	Description  string         `gorm:"not null"`
	Thumbnail    pq.StringArray `gorm:"type:text[]"`
	Link         string         `gorm:"not null"`
	Video        string         `gorm:"not null"`
	Field        string         `gorm:"type:varchar(32);not null"`
	Active       bool           `gorm:"not null"`
	Metadata     string         `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"-"`
}

type ProjectOrmer interface {
	Insert(project Project) (string, error)
	GetOneByCourseworkID(courseworkID string) (Project, error)
	GetMany(skip int, limit int, courseID, name, field string, start, end time.Time) ([]Project, error)
	GetManyByTermAndCourseIDSortByVotes(term time.Time, courseID string) ([]Project, error)
	GetManyByUserIDJoinVote(userID string) ([]Project, error)
	GetManyBookmarkByUserID(userID string) ([]Project, error)
	Update(project Project) error
	UpdateThumbnail(project Project) error
	DeleteByID(courseworkID string) error
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

func (o *projectOrm) GetMany(skip, limit int, courseID, name, field string, startTerm, endTerm time.Time) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN courseworks ON projects.coursework_id = courseworks.id").
		Where("projects.created_at >= ? AND projects.created_at < ?", startTerm, endTerm)
	if courseID != "" {
		result = result.Where("courseworks.course_id = ?", courseID)
	}
	if name != "" {
		result = result.Where("LOWER(projects.name) LIKE LOWER(?)", "%"+name+"%")
	}
	if field != "" {
		result = result.Where(Project{Field: field})
	}
	if limit > 0 {
		result = result.Limit(limit)
	}
	result = result.Offset(skip).Preload("Coursework").Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) GetManyByTermAndCourseIDSortByVotes(term time.Time, courseID string) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN courseworks ON projects.coursework_id = courseworks.id LEFT JOIN votes ON courseworks.id = votes.coursework_id").
		Where("projects.created_at >= ? AND projects.created_at < ? AND courseworks.course_id = ?", utils.TimeToTermTime(term), utils.NextTermTime(term), courseID).
		Group("projects.coursework_id").
		Having("COUNT(votes.id) >= 1").
		Order("COUNT(votes.id) DESC").
		Preload("Coursework").
		Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) GetManyByUserIDJoinVote(userID string) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN votes ON projects.coursework_id = votes.coursework_id").
		Where("votes.user_id = ?", userID).
		Preload("Coursework").
		Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) GetManyBookmarkByUserID(userID string) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN bookmarks ON projects.coursework_id = bookmarks.coursework_id").
		Where("bookmarks.user_id = ?", userID).
		Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) Update(project Project) error {
	// https://gorm.io/docs/update.html#Update-Selected-Fields
	result := o.db.Model(&Project{}).Where("coursework_id = ?", project.CourseworkID).Select("*").Omit("thumbnail", "created_at").Updates(project)
	return result.Error
}

func (o *projectOrm) UpdateThumbnail(project Project) error {
	// https://gorm.io/docs/update.html#Update-Selected-Fields
	result := o.db.Model(&Project{}).Where("coursework_id = ?", project.CourseworkID).Select("thumbnail").Updates(project)
	return result.Error
}

func (o *projectOrm) DeleteByID(courseworkID string) error {
	result := o.db.Model(&Project{}).Where("coursework_id = ?", courseworkID).Delete(&Project{})
	return result.Error
}
