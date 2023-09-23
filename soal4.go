package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ID data yang akan dihapus (gantilah dengan ID yang sesuai)
	postID := 1

	// Buat permintaan DELETE ke API
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID), nil)
	if err != nil {
		fmt.Println("Gagal membuat permintaan DELETE:", err)
		return
	}

	// Kirim permintaan DELETE ke server
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Gagal mengirim permintaan DELETE:", err)
		return
	}
	defer response.Body.Close()

	// Periksa status respons
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Data dengan ID %d berhasil dihapus dengan status code: %d\n", postID, response.StatusCode)
	} else {
		fmt.Printf("Gagal menghapus data dengan ID %d, status code: %d\n", postID, response.StatusCode)
	}
}
