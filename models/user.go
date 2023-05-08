package models

import (
	"github.com/maayarosama/Blogging_system/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model

	// Doesn't override the ID key in gorm model
	ID int `json:"id" gorm:"primaryKey" column:"id"`

	// ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password" binding:"required"`

	Quote string `json:"quote"`
}
type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func (d *DB) SignUp(user *User) *User {
	user.HashedPassword, _ = utils.HashPassword(user.Password)
	d.db.Create(&user)
	return user
}

func (d *DB) GetUsers() []User {
	var Users []User
	d.db.Find(&Users)
	return Users
}
func (d *DB) GetUserByEmail(email string) (*User, error) {
	// Something is wrong here
	var getUserByEmail User
	db := d.db.Where("email=?", email).Find(&getUserByEmail)
	print(db.Error)
	if d.db.Error == nil {
		print("db.Error")
	}

	return &getUserByEmail, d.db.Error
}

// GetUserByid
