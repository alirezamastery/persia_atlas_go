package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"-"`
	Password  string     `gorm:"size:128;not null" json:"-"`
	Mobile    string     `gorm:"size:15;unique" json:"mobile"`
	IsActive  bool       `gorm:"default:true" json:"is_active"`
	IsAdmin   bool       `gorm:"default:false" json:"is_admin"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	Profile   Profile
}

type Profile struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	UserID    uint   `json:"-"`
	FirstName string `gorm:"size:256;default:''" json:"first_name"`
	LastName  string `gorm:"size:256;default:''" json:"last_name"`
	Avatar    string `gorm:"type:text;default:''" json:"avatar"`
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	profile := Profile{UserID: u.ID}
	result := tx.Create(&profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
