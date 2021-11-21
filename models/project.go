package models

import (
	"time"

	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
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
	GetManyByCourseIDAndTerm(courseID string, term, maxTerm time.Time) ([]Project, error)
	GetManyByTermAndCourseIdSortByVotes(term time.Time, courseId string) ([]Project, error)
	Insert(project Project) (courseworkID string, err error)
	UpdateThumbnail(courseworkID string, path string) (err error)
	GetManyBookmarkByUserID(userID string) ([]Project, error)
	GetManyByUserIDJoinVote(userID string) ([]Project, error)
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

func (o *projectOrm) GetManyByTermAndCourseIdSortByVotes(term time.Time, courseId string) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN courseworks ON projects.coursework_id = courseworks.id LEFT JOIN votes ON courseworks.id = votes.coursework_id").
		Where("projects.created_at >= ? AND projects.created_at < ? AND courseworks.course_id = ?", utils.TimeToTermTime(term), utils.NextTermTime(term), courseId).
		Order("Count(votes.id) DESC").
		Group("projects.coursework_id").
		Preload("Coursework").
		Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) GetManyBookmarkByUserID(userID string) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN bookmarks ON projects.coursework_id = bookmarks.coursework_id").
		Where("bookmarks.user_id >= ?", userID).
		Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) GetManyByCourseIDAndTerm(courseID string, term, maxTerm time.Time) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN courseworks ON projects.coursework_id = courseworks.id").
		Where("courseworks.course_id = ? AND projects.created_at >= ? AND projects.created_at < ?", courseID, term, maxTerm).
		Preload("Coursework").
		Find(&projects)
	return projects, result.Error
}

func (o *projectOrm) GetManyByUserIDJoinVote(userID string) ([]Project, error) {
	var projects []Project
	result := o.db.
		Model(&Project{}).
		Joins("INNER JOIN courseworks ON projects.coursework_id = courseworks.id INNER JOIN votes ON courseworks.id = votes.coursework_id").
		Where("votes.user_id = ?", userID).
		Preload("Coursework").
		Find(&projects)
	return projects, result.Error
}
