package models

import (
	"time"

	"github.com/maayarosama/Blogging_system/internal"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id"`
	Name             string    `gorm:"type:varchar(255);not null"`
	Email            string    `gorm:"uniqueIndex;not null"`
	Password         string    `gorm:"not null"`
	Quote            string    `gorm:"type:varchar(255);not null"`
	VerificationCode int
	Verified         bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

func (d *DB) CreateUser(user *User) *User {
	user.ID, _ = uuid.NewUUID()
	user.Password, _ = internal.HashPassword(user.Password)
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
