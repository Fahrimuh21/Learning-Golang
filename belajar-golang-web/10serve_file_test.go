package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	//untuk dynamic file, kita bisa menggunakan http.ServeFile untuk melayani
	// file statis yang ada di dalam folder resources. Dalam contoh ini,
	// jika query parameter name tidak kosong, maka kita akan melayani file ok.html,
	// jika query parameter name kosong, maka kita akan melayani file notfound.html.
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
		//http.ServeFile digunakan untuk melayani file statis
		//  yang ada di dalam folder resources. Dalam contoh ini,
		// jika query parameter name tidak kosong, maka kita akan melayani file ok.html,
		// jika query parameter name kosong, maka kita akan melayani file notfound.html.
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
