package main

import (
	"log"
	"net/http"

	"github.com/gustavocd/file-upload-ajax/controllers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/upload", controllers.UploadFile)

	log.Println("🚀 server running at port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
