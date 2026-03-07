package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
}

//go run 4flag.go -username=eko -password="rahasia banget"
// -host=123.231.23.3 -port=550

// ======================================
// PACKAGE flag - GOLANG
// ======================================
//
// Package flag digunakan untuk:
// - Mengambil argument dari command line
// - Mengatur option / flag (seperti --port, --name)
// - Membuat program CLI lebih rapi & terstruktur
//
// Dibanding os.Args:
// - flag lebih aman
// - otomatis parsing
// - mendukung default value & help
//
// ======================================
// 1. FLAG DASAR
// ======================================
//
// flag.String(name, default, usage)
// flag.Int(name, default, usage)
// flag.Bool(name, default, usage)
//
// Contoh:
// name := flag.String("name", "guest", "nama user")
// port := flag.Int("port", 8080, "port server")
// debug := flag.Bool("debug", false, "mode debug")
//
// ======================================
// 2. flag.Parse()
// ======================================
//
// flag.Parse() WAJIB dipanggil
// agar flag diproses dari command line
//
// Tanpa flag.Parse():
// - semua flag bernilai default
//
// ======================================
// 3. CARA PAKAI FLAG
// ======================================
//
// Jalankan program:
// go run main.go --name=Fahri --port=3000 --debug
//
// Akses nilai flag:
// *name   -> "Fahri"
// *port   -> 3000
// *debug  -> true
//
// Catatan:
// flag mengembalikan POINTER
// harus diakses dengan (*)
//
// ======================================
// 4. FLAG VARIANTS (Var)
// ======================================
//
// Bisa pakai variabel biasa:
//
// var mode string
// flag.StringVar(&mode, "mode", "dev", "mode aplikasi")
//
// Setelah flag.Parse():
// mode -> "dev" / nilai dari CLI
//
// ======================================
// 5. FLAG SHORT vs LONG
// ======================================
//
// Package flag TIDAK mendukung short flag (-p)
// secara otomatis
//
// Harus dibuat manual:
//
// flag.Int("p", 8080, "alias port")
//
// ======================================
// 6. FLAG HELP
// ======================================
//
// Otomatis tersedia:
// -h atau --help
//
// Menampilkan semua flag & usage
//
// ======================================
// 7. SISA ARGUMENT (NON-FLAG)
// ======================================
//
// flag.Args()
// - Mengambil argumen selain flag
//
// flag.NArg()
// - Jumlah argumen non-flag
//
// flag.Arg(i)
// - Argumen ke-i
//
// Contoh:
// go run main.go --name=Fahri file1.txt file2.txt
//
// flag.Args() -> ["file1.txt", "file2.txt"]
//
// ======================================
// 8. FLAG VALIDATION
// ======================================
//
// Validasi manual setelah flag.Parse()
//
// Contoh:
// if *port < 1 || *port > 65535 {
//     log.Fatal("port tidak valid")
// }
//
// ======================================
// 9. CUSTOM USAGE
// ======================================
//
// flag.Usage = func() {
//     fmt.Println("Cara pakai program:")
//     flag.PrintDefaults()
// }
//
// ======================================
// 10. RINGKASAN
// ======================================
//
// flag.String / Int / Bool -> definisi flag
// flag.Parse()            -> parsing flag
// *flagValue              -> ambil nilai
// flag.Args()             -> argumen sisa
// flag.Usage              -> help custom
//
// Gunakan package flag untuk:
// - CLI tool
// - konfigurasi program
// - pengganti os.Args
//
