// Package untuk test file upload
package belajar_golang_web

import (
	"bytes"             // Package untuk membuat buffer (array of bytes)
	_ "embed"           // Package embed untuk embed file (underscore = import tapi tidak digunakan di code, hanya untuk side effect)
	"fmt"               // Package untuk print/format output
	"io"                // Package untuk input/output operations
	"mime/multipart"    // Package untuk handle multipart form data
	"net/http"          // Package untuk HTTP protocol
	"net/http/httptest" // Package untuk test HTTP server
	"os"                // Package untuk operating system operations
	"testing"           // Package untuk unit testing (required untuk *testing.T)
)

// UploadForm adalah handler untuk menampilkan form upload file
func UploadForm(writer http.ResponseWriter, request *http.Request) {
	// Execute template "upload.form.gohtml" dan tulis ke response writer, tidak pass data apapun (nil)
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

// Upload adalah handler untuk memproses file upload
func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(32 << 20) - ini untuk parsing multipart form dengan max size 32MB (tidak dipakai, sudah otomatis)

	// Ambil file dari form field "file", juga dapatkan file header dan error
	file, fileHeader, err := request.FormFile("file")
	// Jika ada error saat mengambil file, panic
	if err != nil {
		panic(err)
	}

	// Buat file di folder ./resources/ dengan nama file yang sama seperti file yang di-upload
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	// Jika error saat create file, panic
	if err != nil {
		panic(err)
	}

	// Copy isi file yang di-upload ke file destination yang baru dibuat
	_, err = io.Copy(fileDestination, file)
	// Jika error saat copy, panic
	if err != nil {
		panic(err)
	}

	// Ambil value dari form field "name"
	name := request.PostFormValue("name")

	// Execute template "upload.success.gohtml" dan pass data berupa map dengan key "Name" dan "File"
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,                             // Tampilkan nama yang diinput di form
		"File": "/static/" + fileHeader.Filename, // Tampilkan URL/path file yang sudah diupload
	})
}

// TestUploadForm adalah test function untuk menjalankan HTTP server dan handle upload form
func TestUploadForm(t *testing.T) {
	// Buat router/mux baru untuk handle HTTP requests
	mux := http.NewServeMux()

	// Register handler UploadForm ke path "/" (root)
	mux.HandleFunc("/", UploadForm)

	// Register handler Upload ke path "/upload"
	mux.HandleFunc("/upload", Upload)

	// Register handler untuk serve static files dari folder ./resources/ di path /static/
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	// Buat HTTP server dengan address localhost:8080 dan handler mux
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// Jalankan server dan listen untuk incoming requests
	err := server.ListenAndServe()
	// Jika ada error, panic
	if err != nil {
		panic(err)
	}
}

// Directive untuk embed file "resources/PNZ-ICON.png" ke dalam binary program
// Variable yang isinya adalah binary data/bytes dari file yang di-embed
//
//go:embed resources/PNZ-ICON.png
var uploadFileTest []byte

// TestUploadFile adalah test function untuk test upload file functionality
func TestUploadFile(t *testing.T) {
	// Buat buffer kosong untuk menyimpan request body
	body := new(bytes.Buffer)

	// Buat multipart writer untuk format multipart/form-data
	writer := multipart.NewWriter(body)

	// Tulis form field "name" dengan value "Eko Kurniawan Khannedy"
	writer.WriteField("name", "Eko Kurniawan Khannedy")

	// Buat form file field "file" dengan filename "CONTOHUPLOAD.png"
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png")

	// Tulis isi file (binary data dari embedded file) ke form file field
	file.Write(uploadFileTest)

	// Close writer untuk finalize/complete multipart data
	writer.Close()

	// Buat HTTP request POST ke "/upload" dengan request body berisi multipart data
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)

	// Set Content-Type header ke multipart/form-data dengan boundary yang sesuai
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Buat response recorder untuk capture response dari handler
	recorder := httptest.NewRecorder()

	// Panggil Upload handler dengan request dan recorder
	Upload(recorder, request)

	// Baca response body dari recorder
	bodyResponse, _ := io.ReadAll(recorder.Result().Body)

	// Print response body ke console
	fmt.Println(string(bodyResponse))
}
