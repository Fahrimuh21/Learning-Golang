// Paket belajar_golang_web untuk mempelajari fitur download file di Go
package belajar_golang_web

// Import library yang dibutuhkan
import (
	// untuk formatting output
	"fmt"
	// untuk HTTP server dan client
	"net/http"
	// untuk testing framework
	"testing"
)

// Fungsi DownloadFile untuk handle request download file dari client
// Parameter: writer untuk response, request untuk request dari client
func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	// Ambil parameter "file" dari URL query
	file := request.URL.Query().Get("file")

	// Cek apakah parameter file kosong
	if file == "" {
		// Set response status code menjadi Bad Request (400)
		writer.WriteHeader(http.StatusBadRequest)
		// Tulis pesan error ke response
		fmt.Fprint(writer, "Bad Request")
		// Return untuk menghentikan eksekusi fungsi
		return
	}

	// Set header Content-Disposition untuk download dengan nama file tertentu
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	// Serve file dari folder resources sesuai dengan nama file yang diminta
	http.ServeFile(writer, request, "./resources/"+file)
}

// Fungsi TestDownloadFile untuk test fitur download file menggunakan http server
// Parameter: t adalah testing.T untuk testing utilities
func TestDownloadFile(t *testing.T) {
	// Buat instance HTTP server baru
	server := http.Server{
		// Set alamat server: localhost port 8080
		Addr: "localhost:8080",
		// Set handler untuk menangani request ke DownloadFile function
		Handler: http.HandlerFunc(DownloadFile),
	}

	// Jalankan server dan dengarkan request dari client
	err := server.ListenAndServe()
	// Cek apakah ada error saat menjalankan server
	if err != nil {
		// Jika ada error, panic untuk menghentikan program
		panic(err)
	}
}
