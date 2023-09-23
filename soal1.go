package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Kirim permintaan GET ke API
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Gagal mengirim permintaan GET:", err)
		return
	}
	defer response.Body.Close()

	// Dekode respons JSON
	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		fmt.Println("Gagal mendekode respons JSON:", err)
		return
	}

	// Tampilkan data
	for _, post := range posts {
		fmt.Printf("ID: %d\nUserID: %d\nTitle: %s\n\n", post.ID, post.UserID, post.Title)
	}
}
