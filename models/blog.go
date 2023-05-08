// Ignore this file for now as it needs a lot of modifications
package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ID      int    `json:"id" gorm:"primaryKey"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (blog *Blog) CreateBlog() *Blog {
	// db.NewRecord(user)
	db.Create(&blog)
	return blog
}

func GetBlogs() []Blog {
	var Blogs []Blog
	db.Find(&Blogs)
	return Blogs
}

func GetBlogById(id int64) (*Blog, *gorm.DB) {
	var getBlog Blog
	db := db.Where("blogId=?", id).Find(&getBlog)
	return &getBlog, db
}

func DeleteBlog(id int64) Blog {
	var blog Blog
	db.Where("blogId=?", id).Delete(blog)
	return blog
}
