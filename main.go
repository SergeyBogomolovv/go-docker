package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	h := &handler{db: &db{posts: make([]Post, 0)}}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/posts", h.GetPosts)
	r.Post("/posts", h.AddPost)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Printf("Listening on %s\n", addr)
	srv.ListenAndServe()
}

type handler struct {
	db *db
}

func (h *handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := h.db.GetPosts()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (h *handler) AddPost(w http.ResponseWriter, r *http.Request) {
	var p Post
	json.NewDecoder(r.Body).Decode(&p)
	h.db.AddPost(p)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type db struct {
	posts []Post
}

func (d *db) GetPosts() []Post {
	return d.posts
}

func (d *db) AddPost(p Post) {
	d.posts = append(d.posts, p)
}
