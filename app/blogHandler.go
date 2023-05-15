// Ignore this file for now as it needs a lot of modifications
package app

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	"github.com/maayarosama/Blogging_system/models"
// 	"github.com/maayarosama/Blogging_system/utils"
// )

// // var newBlog models.Blog

// func GetBlogs(w http.ResponseWriter, r *http.Request) {

// 	// blogs:= models.GetBlogs()
// 	newBlogs := models.GetBlogs()
// 	res, _ := json.Marshal(newBlogs)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetBlogById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	blogId := vars["blog_id"]
// 	ID, err := strconv.ParseInt(blogId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error parsing blog")
// 	}
// 	blogDetails, _ := models.GetBlogById(ID)
// 	res, _ := json.Marshal(blogDetails)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

// func CreateBlog(w http.ResponseWriter, r *http.Request) {
// 	createBlog := &models.Blog{}
// 	utils.ParseBody(r, createBlog)
// 	b := createBlog.CreateBlog()
// 	res, _ := json.Marshal(b)

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteBlog(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	blogId := vars["blog_id"]
// 	ID, err := strconv.ParseInt(blogId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error parsing blog")
// 	}

// 	blog := models.DeleteBlog(ID)
// 	res, _ := json.Marshal(blog)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

// // func UpdateBlog(w http.ResponseWriter, r *http.Request) {
// // 	var updateBlog = &models.Blog{}
// // 	utils.ParseBody(r, updateBlog)
// // 	vars := mux.Vars(r)
// // 	blogId := vars["blog_id"]
// // 	ID, err := strconv.ParseInt(blogId, 0, 0)
// // 	if err != nil {
// // 		fmt.Println("error parsing blog")
// // 	}
// // 	blogDetails , db := models.GetBlogById(ID)
// // 	if

// // }
