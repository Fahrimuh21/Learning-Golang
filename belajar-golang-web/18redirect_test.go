package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// redirect itu untuk mengarahkan user ke halaman lain, bisa ke halaman yang ada di dalam aplikasi kita, bisa juga ke halaman yang ada di luar aplikasi kita, seperti google.com, facebook.com, dan lain-lain.
// untuk melakukan redirect, kita bisa menggunakan function http.Redirect, dengan parameter writer, request, url tujuan, dan status code redirect yang ingin kita gunakan, seperti http.StatusTemporaryRedirect, http.StatusMovedPermanently, dan lain-lain.
func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
	//ini akan menampilkan tulisan "Hello Redirect" di halaman tujuan redirect, karena kita melakukan redirect ke halaman yang ada di dalam aplikasi kita, yaitu /redirect-to, sehingga ketika user mengakses /redirect-from, maka user akan diarahkan ke /redirect-to, dan di halaman /
	//redirect-to, maka akan menampilkan tulisan "Hello Redirect".
}

// untuk melakukan redirect ke halaman lain, kita bisa menggunakan function http.Redirect, dengan parameter writer, request, url tujuan, dan status code redirect yang ingin kita gunakan, seperti http.StatusTemporaryRedirect, http.StatusMovedPermanently, dan lain-lain.
func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
	//ke web kita bisa menggunakan url tujuan dengan menggunakan path, seperti "/redirect-to" pada contoh di atas, sehingga ketika user mengakses path "/redirect-from", maka user akan diarahkan ke path "/redirect-to", dan ketika user mengakses path "/redirect-to", maka user akan melihat tampilan yang berisi "Hello Redirect".
}

// untuk melakukan redirect ke halaman lain, kita bisa menggunakan function http.Redirect, dengan parameter writer, request, url tujuan, dan status code redirect yang ingin kita gunakan, seperti http.StatusTemporaryRedirect, http.StatusMovedPermanently, dan lain-lain.
func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "https://www.programmerzamannow.com/", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
