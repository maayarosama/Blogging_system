package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maayarosama/Blogging_system/internal"
	"github.com/maayarosama/Blogging_system/models"
	"github.com/rs/zerolog/log"
)

func (a *App) GetBlogs(w http.ResponseWriter, r *http.Request) {
	blogs := a.db.GetBlogs()
	res, _ := json.Marshal(blogs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func (a *App) GetUsersBlogs(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDKey("UserID")).(int)
	blogs := a.db.GetUsersBlogs(userID)
	res, _ := json.Marshal(blogs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (a *App) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId := vars["BlogId"]
	ID, err := strconv.Atoi(blogId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))
		return
	}
	blogDetails, _ := a.db.GetBlogByID(ID)
	res, _ := json.Marshal(blogDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (a *App) CreateBlog(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDKey("UserID")).(int)
	createBlog := &models.Blog{}
	ParseBody(r, createBlog)
	createBlog.Userid = int(userID)

	b := a.db.CreateBlog(createBlog)
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (a *App) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId := vars["BlogId"]
	ID, err := strconv.Atoi(blogId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))
		return
	}
	deleteBlog, err := a.db.GetBlogByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))

		return
	}
	ParseBody(r, deleteBlog)

	err = a.db.DeleteBlog(deleteBlog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully \n"))

}

func (a *App) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId := vars["BlogId"]
	ID, err := strconv.Atoi(blogId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))
		return
	}
	deleteBlog, err := a.db.GetBlogByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))

		return
	}
	ParseBody(r, deleteBlog)

	err = a.db.DeleteBlog(deleteBlog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("failed Deletation: %v", err)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully \n"))

}
