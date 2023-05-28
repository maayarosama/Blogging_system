package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// type mockDB interface {
// 	CreateUser(user *User) *User
// 	GetUsers() []User
// 	GetUserByEmail(email string) (*User, error)
// 	UpdateUser(user *User) error
// 	UpdateUserVerfied(user *User) error
// 	GetBlogs() []Blog
// 	GetUsersBlogs(id int) []Blog
// 	CreateBlog(b *Blog) *Blog
// 	GetBlogByID(id int) (*Blog, error)
// 	DeleteBlog(b *Blog) error
// }

type DB struct {
	db *gorm.DB
}

// func NewDB() mockDB {
// 	return DB{}

// }

// NewDB creates new DB
func NewDB() DB {
	return DB{}
}

// Connect connects to database file
func (d *DB) Connect(path string) error {
	gormDB, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = gormDB
	return nil
}

// Migrate migrates db schema
func (d *DB) Migrate() error {
	err := d.db.AutoMigrate(&User{}, &Blog{})
	if err != nil {
		return err
	}
	return nil

}
