package main

import (
	"encoding/json"
	"net/http"
	"math/rand"

	"./entity"
	"./repository"
)


var (
	repo repository.PostRepository = repository.NewRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	post.ID = rand.Int()
	repo.Save(post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}