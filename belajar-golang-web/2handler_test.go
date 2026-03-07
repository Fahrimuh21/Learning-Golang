package belajar_golang_web

//handler adalah fungsi yang akan dipanggil ketika ada request yang masuk ke server. Handler menerima dua parameter, yaitu http.ResponseWriter dan *http.Request.
//http.ResponseWriter adalah interface yang digunakan untuk menulis response ke client. Dengan http.ResponseWriter, kita bisa menulis header, status code, dan body dari response yang akan dikirim ke client.
//*http.Request adalah struct yang berisi informasi tentang request yang masuk, seperti method, URL, header, dan body dari request tersebut.

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//handler adalah fungsi yang akan dipanggil ketika ada request yang masuk ke server. Handler menerima dua parameter, yaitu http.ResponseWriter dan *http.Request.
		//http.ResponseWriter adalah interface yang digunakan untuk menulis response ke client. Dengan http.ResponseWriter, kita bisa menulis header, status code, dan body dari response yang akan dikirim ke client.
		//*http.Request adalah struct yang berisi informasi tentang request yang masuk, seperti method, URL, header, dan body dari request tersebut.
		// logic web
		fmt.Fprint(writer, "Hello World")
		//fmt.Fprint digunakan untuk menulis string ke http.ResponseWriter. Dalam contoh ini, kita menulis "Hello World" ke http.ResponseWriter, sehingga ketika ada request yang masuk ke server, client akan menerima response dengan body "Hello World".
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
		//masukan handler ke dalam server, sehingga ketika ada request yang masuk ke server, handler akan dipanggil untuk menangani request tersebut.
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Image")
	})
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
