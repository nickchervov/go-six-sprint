package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	resp, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if file == nil {
		http.Error(w, "must upload a file", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, "file upload error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading the uploaded file", http.StatusInternalServerError)
		return
	}

	convertedData, err := service.ConvertingData(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	timeStamp := time.Now().Format("02-01-06_15-04-05")
	fileExt := filepath.Ext(handler.Filename)
	localFileName := filepath.Join(timeStamp, fileExt)

	if err := os.WriteFile(localFileName, []byte(convertedData), 0755); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedData))
}
