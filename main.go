package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustavocd/file-upload-ajax/controllers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/upload", controllers.UploadFile)

	fmt.Println("Server running")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
