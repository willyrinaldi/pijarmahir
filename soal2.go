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
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/3")
	if err != nil {
		fmt.Println("Gagal mengirim permintaan GET:", err)
		return
	}
	defer response.Body.Close()

	// Memeriksa apakah data dengan ID 3 ditemukan
	if response.StatusCode == http.StatusNotFound {
		fmt.Println("Data dengan ID 3 tidak ditemukan.")
		return
	}

	// Dekode respons JSON
	var post Post
	err = json.NewDecoder(response.Body).Decode(&post)
	if err != nil {
		fmt.Println("Gagal mendekode respons JSON:", err)
		return
	}

	// Tampilkan data
	fmt.Printf("ID: %d\nUserID: %d\nTitle: %s\n\n", post.ID, post.UserID, post.Title)
}
