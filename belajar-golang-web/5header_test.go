package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type") //request.Header.Get digunakan untuk
	// mendapatkan nilai dari header yang dikirimkan oleh client. Dalam contoh ini, kita
	//  mencoba untuk mendapatkan nilai dari header Content-Type yang dikirimkan oleh client.
	// Jika header Content-Type tidak ada, maka Get akan mengembalikan string kosong ("").
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	//request.Header digunakan untuk mengakses header dari request yang masuk.
	// Dengan request.Header, kita bisa membaca header yang dikirimkan oleh client,
	// seperti Content-Type, User-Agent, Authorization, dan lain-lain. Dalam

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Programmer Zaman Now")
	//writer.Header().Add digunakan untuk menambahkan header
	// ke response yang akan dikirimkan ke client. Dalam
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	fmt.Println(response.Header.Get("x-powered-by"))
}
