package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFile uploads a file to the server
func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err.Error())
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg", "image/jpg", "image/png":
		err := saveFile(w, file, handle)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "No se pudo guardar el archivo")
			return
		}
		jsonResponse(w, http.StatusCreated, "Archivo guardado con éxito")
	default:
		jsonResponse(w, http.StatusBadRequest, "El formato de la imagen no es válido")
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) error {
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = os.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
