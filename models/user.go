package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(32);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	IsAdmin   bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

type UserOrmer interface {
	Insert(user User) (id string, err error)
	GetOneByEmail(email string) (user User, err error)
	Update(user User) (err error)
}

type userOrm struct {
	db *gorm.DB
}

func NewUserOrmer(db *gorm.DB) UserOrmer {
	_ = db.AutoMigrate(&User{})
	return &userOrm{db}
}

func (o *userOrm) Insert(user User) (id string, err error) {
	result := o.db.Model(&User{}).Create(&user)
	return user.ID, result.Error
}

func (o *userOrm) GetOneByEmail(email string) (user User, err error) {
	result := o.db.Model(&User{}).Where("email = ?", email).First(&user)
	return user, result.Error
}

func (o *userOrm) Update(user User) (err error) {
	result := o.db.Model(&User{}).Where("email = ?", user.Email).Select("is_admin").Updates(&user)
	return result.Error
}
