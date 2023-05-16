package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ID        int       `gorm:"primary_key; unique; AUTO_INCREMENT; column:id"`
	Userid    int       `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (d *DB) GetBlogs() []Blog {
	var Blogs []Blog
	d.db.Find(&Blogs)
	return Blogs
}

func (d *DB) GetUsersBlogs(id int) []Blog {
	var Blogs []Blog

	d.db.Where("userid = ?", id).Find(&Blogs)
	return Blogs
}
func (d *DB) CreateBlog(b *Blog) *Blog {
	d.db.Create(&b)
	return b
}

func (d *DB) GetBlogByID(id int) (*Blog, error) {
	var b Blog
	res := d.db.First(&b, "id = ?", id)
	return &b, res.Error
}

func (d *DB) DeleteBlog(b *Blog) error {
	query := d.db.Where("id=?", b.ID).Delete(b)
	return query.Error
}
