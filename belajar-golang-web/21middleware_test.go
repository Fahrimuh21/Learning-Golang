// Package belajar_golang_web berisi implementasi middleware HTTP dan contoh penggunaannya
package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// LogMiddleware adalah middleware yang berfungsi untuk mencatat (logging) setiap request
// yang masuk sebelum dan sesudah handler dieksekusi.
// Middleware ini mengimplementasikan interface http.Handler melalui method ServeHTTP
type LogMiddleware struct {
	// Handler adalah handler HTTP berikutnya yang akan dieksekusi dalam chain
	Handler http.Handler
}

// ServeHTTP mengimplementasikan interface http.Handler untuk LogMiddleware.
// Method ini mencatat output sebelum dan sesudah handler dieksekusi sehingga dapat
// melacak alur eksekusi request.
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

// ErrorHandler adalah middleware yang berfungsi untuk menangkap panic dan error
// yang mungkin terjadi selama eksekusi handler, sehingga server tidak crash.
// Middleware ini mengimplementasikan interface http.Handler melalui method ServeHTTP
type ErrorHandler struct {
	// Handler adalah handler HTTP berikutnya yang akan dieksekusi dalam chain
	Handler http.Handler
}

// ServeHTTP mengimplementasikan interface http.Handler untuk ErrorHandler.
// Method ini menggunakan defer dan recover untuk menangkap panic yang terjadi
// dan mengembalikan response error dengan status code 500 kepada client.
func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

// TestMiddleware adalah test case yang mendemonstrasikan penggunaan middleware chaining
// dalam Go. Test ini membuat server HTTP dengan kombinasi LogMiddleware dan ErrorHandler
// untuk menangani request dan error secara terpadu.
// Server menyediakan 3 endpoint: "/" (normal), "/foo" (normal), dan "/panic" (panic handler)
//
// ALUR EKSEKUSI MIDDLEWARE CHAIN:
// Ketika request masuk ke server, alur eksekusinya adalah:
// 1. ErrorHandler.ServeHTTP() dipanggil (middleware outer layer - error handling)
// 2. → LogMiddleware.ServeHTTP() dipanggil (middleware inner layer - logging)
// 3. → mux.ServeHTTP() dipanggil (router yang mengarahkan ke handler spesifik)
// 4. → Handler endpoint (~/ /foo /panic) dieksekusi
// 5. ← Kembali ke LogMiddleware (cetak "After Execute Handler")
// 6. ← Kembali ke ErrorHandler (jika ada panic, tangkap di sini)
func TestMiddleware(t *testing.T) {
	// Step 1: Buat router HTTP menggunakan ServeMux
	mux := http.NewServeMux()

	// Daftarkan endpoint "/" - mengembalikan "Hello Middleware"
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})

	// Daftarkan endpoint "/foo" - mengembalikan "Hello Foo"
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})

	// Daftarkan endpoint "/panic" - memicu panic untuk ditest error handling
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		panic("Ups")
	})

	// Step 2: Bungkus mux dengan LogMiddleware (layer pertama middleware)
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	// Step 3: Bungkus logMiddleware dengan ErrorHandler (layer kedua/outer middleware)
	// Dengan urutan ini, setiap request akan melewati ErrorHandler → LogMiddleware → mux
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	// Step 4: Buat HTTP server yang mendengarkan port 8080
	// Handler: errorHandler adalah entry point pertama saat request masuk
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	// Step 5: JALANKAN SERVER - ListenAndServe() akan memblok eksekusi
	// Server akan terus berjalan di background mendengarkan request pada localhost:8080
	// Ketika test function dijalankan (biasanya dengan: go test),
	// server akan mulai listening dan siap menerima HTTP request
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
