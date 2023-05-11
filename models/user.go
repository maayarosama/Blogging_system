package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/maayarosama/Blogging_system/utils"
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

// type User struct {
// 	gorm.Model

// 	// Doesn't override the ID key in gorm model
// 	ID int `json:"id" gorm:"primaryKey" column:"id"`

// 	// ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	// HashedPassword string `json:"hashed_password" binding:"required"`

//		Quote string `json:"quote"`
//	}

func (d *DB) CreateUser(user *User) *User {
	user.ID, _ = uuid.NewUUID()
	user.Password, _ = utils.HashPassword(user.Password)
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
	println(res)
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
