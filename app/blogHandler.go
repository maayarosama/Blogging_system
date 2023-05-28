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
	blogs := a.DB.GetBlogs()
	res, err := json.Marshal(blogs)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

}
func (a *App) GetUsersBlogs(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDKey("UserID")).(int)
	blogs := a.DB.GetUsersBlogs(userID)
	res, err := json.Marshal(blogs)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
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
	blogDetails, _ := a.DB.GetBlogByID(ID)
	res, _ := json.Marshal(blogDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

}

func (a *App) CreateBlog(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(internal.UserIDKey("UserID")).(int)
	createBlog := &models.Blog{}
	ParseBody(r, createBlog)
	createBlog.Userid = userID

	b := a.DB.CreateBlog(createBlog)
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
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
	deleteBlog, err := a.DB.GetBlogByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))

		return
	}
	ParseBody(r, deleteBlog)

	err = a.DB.DeleteBlog(deleteBlog)
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
	deleteBlog, err := a.DB.GetBlogByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("error: %v", err)
		w.Write([]byte(err.Error()))

		return
	}
	ParseBody(r, deleteBlog)

	err = a.DB.DeleteBlog(deleteBlog)
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
