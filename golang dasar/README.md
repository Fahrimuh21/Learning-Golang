Oke, aku bikinkan **README.md** untuk pembelajaran **Golang dasar**. Struktur ini cocok kalau kamu mau taruh di GitHub atau dokumentasi belajar pribadi 👇

---

````markdown
# 🚀 Belajar Golang Dasar

Golang (atau Go) adalah bahasa pemrograman open-source yang dikembangkan oleh Google.  
Bahasa ini dirancang agar **sederhana, cepat, aman, dan efisien**, terutama untuk aplikasi skala besar dan sistem terdistribusi.

---

## 📌 Daftar Isi
1. [Pengantar Golang](#-pengantar-golang)
2. [Instalasi](#-instalasi)
3. [Program Pertama](#-program-pertama)
4. [Dasar-Dasar Golang](#-dasar-dasar-golang)
   - Variabel & Konstanta
   - Tipe Data
   - Operator
   - Percabangan (`if`, `switch`)
   - Perulangan (`for`)
   - Array, Slice, dan Map
   - Fungsi (termasuk variadic)
5. [Concurrency: Goroutine & Channel](#-concurrency-goroutine--channel)
6. [Error Handling](#-error-handling)
7. [Contoh Project Sederhana](#-contoh-project-sederhana)
8. [Referensi](#-referensi)

---

## 🟢 Pengantar Golang
- **Go** adalah bahasa pemrograman yang:
  - **Statically typed** → tipe data harus jelas
  - **Compiled** → menghasilkan binary yang cepat
  - Mendukung **concurrency** dengan `goroutine`
  - Memiliki **garbage collector** bawaan

---

## ⚙️ Instalasi
1. Download di [https://go.dev/dl/](https://go.dev/dl/)  
2. Cek versi:
   ```bash
   go version
````

3. Setup workspace (misalnya di folder `go-projects`).

---

## 👨‍💻 Program Pertama

Buat file `main.go`:

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, Golang!")
}
```

Jalankan dengan:

```bash
go run main.go
```

---

## 📖 Dasar-Dasar Golang

### 1. Variabel & Konstanta

```go
var name string = "Fahri"
age := 20          // short declaration
const PI = 3.14    // konstanta
```

### 2. Tipe Data

* `string`, `int`, `float64`, `bool`
* Struktur data: `array`, `slice`, `map`

### 3. Percabangan

```go
if age >= 18 {
    fmt.Println("Dewasa")
} else {
    fmt.Println("Anak-anak")
}
```

### 4. Perulangan

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 5. Array, Slice, dan Map

```go
arr := [3]int{1, 2, 3}
slice := []string{"go", "python", "java"}
m := map[string]int{"a": 1, "b": 2}
```

### 6. Fungsi

```go
func tambah(a int, b int) int {
    return a + b
}
```

#### Variadic Function

```go
func jumlahkan(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

---

## ⚡ Concurrency: Goroutine & Channel

```go
go sayHello() // menjalankan fungsi secara paralel

ch := make(chan string)
go func() {
    ch <- "Halo dari goroutine"
}()
fmt.Println(<-ch)
```

---

## 🚨 Error Handling

```go
import (
    "errors"
    "fmt"
)

func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("tidak bisa membagi dengan nol")
    }
    return a / b, nil
}
```

---

## 📂 Contoh Project Sederhana

* CLI kalkulator
* Rest API dengan `net/http`
* Aplikasi goroutine sederhana

---

## 📚 Referensi

* [Dokumentasi Resmi Go](https://go.dev/doc/)
* [Go by Example](https://gobyexample.com/)
* [Tour of Go](https://go.dev/tour/)

---

```

---

Mau aku bikinkan juga versi **contoh project kecil (misalnya kalkulator CLI atau REST API sederhana)** untuk ditaruh di README ini biar lebih menarik?
```
