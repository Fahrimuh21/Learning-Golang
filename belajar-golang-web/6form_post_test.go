package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Eko&last_name=Khannedy")
	//strings.NewReader digunakan untuk membuat io.Reader dari string yang kita buat.
	// Dalam contoh ini, kita membuat io.Reader dari string "first_name=Eko&last_name=Khannedy",
	// sehingga ketika kita membuat request dengan body requestBody, maka body dari
	// request tersebut akan berisi "first_name=Eko&last_name=Khannedy".
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//udah spesifikasi content type, karena form post biasanya menggunakan
	// content type application/x-www-form-urlencoded, sehingga server bisa
	// memproses data yang dikirimkan dengan benar.

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
