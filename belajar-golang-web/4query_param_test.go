package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	// Jika query parameter name tidak ada, maka akan menampilkan "Hello" saja
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Eko", nil)
	//query parameter = data yang dikirimkan melalui URL setelah tanda tanya (?)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Eko&last_name=Khannedy", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
	// query["name"] digunakan untuk mendapatkan semua nilai dari query parameter name yang dikirimkan melalui URL. Karena query parameter name bisa memiliki lebih dari satu nilai (seperti pada contoh URL http://localhost:8080/hello?name=Eko&name=Kurniawan&name=Khannedy), maka query["name"] akan mengembalikan slice ([]string) yang berisi semua nilai dari query parameter name.
	// strings.Join digunakan untuk menggabungkan semua nilai dari query parameter name menjadi satu string dengan spasi sebagai pemisah. Dengan cara ini, jika kita mengirimkan URL http://localhost:8080/hello?name=Eko&name=Kurniawan&name=Khannedy, maka output yang dihasilkan akan menjadi "Hello Eko Kurniawan Khannedy".
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Eko&name=Kurniawan&name=Khannedy", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
