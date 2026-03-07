package belajar_golang_web

import (
	"net/http" //package http adalah package yang digunakan untuk membuat web server dan client HTTP
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{ //ini adalah struct yang digunakan untuk mengkonfigurasi dan menjalankan server HTTP. Struct ini memiliki beberapa field yang dapat diatur, seperti Addr, Handler, TLSConfig, dan lain-lain.
		Addr: "localhost:8080",
		//rekomendasi 4 digit untuk port, karena 2 digit pertama bisa digunakan untuk development, dan 2 digit terakhir bisa digunakan untuk production
	}

	err := server.ListenAndServe()
	//listenAndServe akan mengembalikan error ketika server berhenti, baik karena error atau karena server dimatikan dengan benar.
	//Jika server berhenti karena error, maka err akan berisi informasi tentang error tersebut. Namun, jika server dimatikan dengan benar
	// (misalnya dengan Ctrl+C), ListenAndServe akan mengembalikan error http.ErrServerClosed, yang menunjukkan bahwa server telah ditutup dengan benar.
	if err != nil {
		panic(err)
	}
}
