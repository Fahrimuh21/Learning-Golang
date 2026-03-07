package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

// kaya pengisian handler, tapi kita tidak perlu menjalankan server untuk mengetes handler, kita bisa langsung memanggil handler dengan menggunakan http.ResponseWriter dan *http.Request yang kita buat sendiri. Dengan cara ini, kita bisa mengetes handler dengan lebih mudah dan cepat tanpa harus menjalankan server terlebih dahulu.
func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	//newrequest digunakan untuk membuat request baru yang akan digunakan untuk testing.
	//Dalam contoh ini, kita membuat request dengan method GET, URL http://localhost:8080/hello, dan body nil (karena GET request tidak memiliki body).
	recorder := httptest.NewRecorder()
	//newrecorder digunakan untuk membuat http.ResponseWriter yang akan digunakan untuk merekam response yang dihasilkan oleh handler.
	HelloHandler(recorder, request)
	//memanggil handler dengan menggunakan recorder sebagai http.ResponseWriter dan request sebagai *http.Request. Setelah handler dipanggil, recorder akan berisi response yang dihasilkan oleh handler.
	//tidak menjalankan server, tapi langsung memanggil handler dengan menggunakan http.ResponseWriter dan *http.Request yang kita buat sendiri. Dengan cara ini, kita bisa mengetes handler dengan lebih mudah
	// dan cepat tanpa harus menjalankan server terlebih dahulu.

	response := recorder.Result()
	//result digunakan untuk mendapatkan http.Response dari recorder setelah handler dipanggil. Dengan http.Response ini, kita bisa membaca status code, header, dan body dari response yang dihasilkan oleh handler.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	//io.ReadAll digunakan untuk membaca body dari http.Response. Body dari http.Response adalah io.ReadCloser, sehingga kita perlu menggunakan io.ReadAll untuk membaca seluruh isi body dan mengembalikannya sebagai byte slice ([]byte).
	bodyString := string(body)

	fmt.Println(bodyString)
}
