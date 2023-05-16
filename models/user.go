package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               int    `gorm:"primary_key; unique; AUTO_INCREMENT; column:id"`
	Name             string `gorm:"type:varchar(255);not null"`
	Email            string `gorm:"uniqueIndex;not null"`
	Password         string `gorm:"type:varchar(255);not null"`
	Quote            string `gorm:"type:varchar(255);not null"`
	VerificationCode int
	Verified         bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

func (d *DB) CreateUser(user *User) *User {
	d.db.Create(&user)
	return user
}

func (d *DB) GetUsers() []User {
	var Users []User
	d.db.Find(&Users)
	return Users
}

func (d *DB) GetUserByEmail(email string) (*User, error) {
	var u User
	res := d.db.First(&u, "email = ?", email)
	return &u, res.Error
}
func (d *DB) UpdateUser(user *User) error {
	query := d.db.Model(&User{}).Where("ID = ?", user.ID).Updates(user)
	return query.Error
}
func (d *DB) UpdateUserVerfied(user *User) error {
	query := d.db.Model(&User{}).Where("ID = ?", user.ID).Update("verified", user.Verified)
	return query.Error
}

// GetUserByid
