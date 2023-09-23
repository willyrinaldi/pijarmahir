package main

import (
	"bytes"
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
	// Data postingan yang akan disimpan
	newPost := Post{
		UserID: 1,
		ID:     101, // ID akan di-generate oleh server, jadi Anda bisa mengisi apa saja
		Title:  "New Post Title",
		Body:   "This is the body of the new post.",
	}

	// Mengkonversi data postingan menjadi JSON
	postData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("Gagal mengonversi data ke JSON:", err)
		return
	}

	// Kirim permintaan POST ke API
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("Gagal mengirim permintaan POST:", err)
		return
	}
	defer response.Body.Close()

	// Periksa status respons
	if response.StatusCode == http.StatusCreated {
		fmt.Println("Data postingan berhasil disimpan dengan status code:", response.StatusCode)
	} else {
		fmt.Println("Gagal menyimpan data postingan dengan status code:", response.StatusCode)
	}
}
