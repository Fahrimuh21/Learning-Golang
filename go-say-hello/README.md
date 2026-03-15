# Go Say Hello Module

Repositori ini berisi contoh modul Go yang sederhana untuk pembelajaran tentang Go Modules. Modul ini menyediakan fungsi untuk mengucapkan salam.

## Struktur Folder

```
go-say-hello/
├── go.mod                          # File modul Go
└── say_hello.go                    # Implementasi fungsi SayHello
```

## Deskripsi File

- **go.mod**: Mendefinisikan modul dengan nama `github.com/ProgrammerZamanNow/go-say-hello/v2` dan menggunakan Go versi 1.15.
- **say_hello.go**: Berisi package `go_say_hello` dengan fungsi `SayHello(name string) string` yang mengembalikan string "Hello " + name.

## Cara Menggunakan Modul

Untuk menggunakan modul ini dalam proyek lain:

1. Inisialisasi modul di proyek Anda:
   ```bash
   go mod init nama-modul-anda
   ```

2. Tambahkan dependency:
   ```bash
   go get github.com/ProgrammerZamanNow/go-say-hello/v2
   ```

3. Import dan gunakan dalam kode:
   ```go
   import go_say_hello "github.com/ProgrammerZamanNow/go-say-hello/v2"

   func main() {
       fmt.Println(go_say_hello.SayHello("Nama Anda"))
   }
   ```

## Cara Menjalankan (Jika Dijalankan Langsung)

Meskipun ini adalah modul library, Anda bisa menjalankannya dengan membuat file main.go sementara:

```bash
go run say_hello.go  # Tapi ini akan error karena bukan main package
```

Sebaiknya gunakan dalam aplikasi terpisah seperti di folder `app-say-hello`.

## Modul Go

- **Nama Modul**: `github.com/ProgrammerZamanNow/go-say-hello/v2`
- **Versi Go**: 1.15

## Persyaratan

- Go versi 1.15 atau lebih baru

## Catatan

Ini adalah contoh sederhana untuk mempelajari konsep Go Modules, versioning, dan publishing modul ke GitHub. Modul ini dapat dipublikasikan dan digunakan oleh proyek lain.