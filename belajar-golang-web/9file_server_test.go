package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//ini adalah contoh penggunaan http.FileServer untuk membuat file server yang melayani file s
// tatis dari direktori resources.

func TestFileServer(t *testing.T) {

	directory := http.Dir("./resources")
	//ambil direktori resources, karena kita akan menyimpan file statis di dalam folder resources
	fileServer := http.FileServer(directory)
	//http.FileServer digunakan untuk membuat file server yang akan melayani file-file statis
	//yang ada di dalam direktori resources.

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//jadi ini kita menggunakan http.StripPrefix untuk menghapus prefix /static
	// dari URL yang masuk sebelum diteruskan ke file server.
	// Misalnya, jika ada request ke /static/index.html,
	//maka http.StripPrefix akan menghapus prefix /static sehingga file server akan menerima request untuk /index.html.
	// Dengan cara ini, kita bisa menyimpan file-file statis di dalam folder resources tanpa harus menambahkan prefix /static di dalam folder tersebut.

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed resources
var resources embed.FS

//embed digunakan untuk meng-embed file atau direktori ke dalam binary Go.
//Dalam contoh ini, kita meng-embed direktori resources ke dalam variabel resources yang bertipe embed.FS. Dengan cara ini, kita bisa mengakses file-file yang ada di dalam direktori resources melalui variabel resources tanpa harus menyimpan file-file tersebut di dalam folder resources secara fisik.

//ini adalah contoh penggunaan http.FileServer untuk membuat file server yang melayani file statis dari embed.FS.

func TestFileServerGolangEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	//fs.Sub digunakan untuk mengambil sub direktori
	//dari embed.FS. Dalam contoh ini, kita mengambil
	//sub direktori resources dari embed.FS yang sudah kita buat sebelumnya.
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//static dihapus karena kita sudah menggunakan fs.Sub untuk
	// mengambil sub direktori resources, sehingga file server
	// akan menerima request untuk /index.html langsung tanpa
	// harus menghapus prefix /static lagi.

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
