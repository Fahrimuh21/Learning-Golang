# Belajar Golang Unit Test

Repositori ini berisi materi pembelajaran tentang Unit Testing dalam bahasa pemrograman Go. Folder ini mencakup berbagai konsep testing mulai dari basic test, benchmark, hingga mocking dengan testify.

## Struktur Folder

```
belajar-golang-unit-test/
├── go.mod                          # File modul Go dengan dependency testify
├── .gitignore                      # File gitignore
├── entity/                         # Package entity
│   └── category.go                 # Struct Category
├── helper/                         # Package helper
│   ├── hello_world.go              # Fungsi HelloWorld
│   └── hello_world_test.go         # Unit test dan benchmark untuk HelloWorld
├── repository/                     # Package repository
│   ├── category_repository.go      # Interface CategoryRepository
│   └── category_repository_mock.go # Mock implementation untuk testing
└── service/                        # Package service
    ├── category_service.go         # Service yang menggunakan repository
    └── category_service_test.go    # Unit test untuk service dengan mock
```

## Deskripsi File

### Entity
- **category.go**: Mendefinisikan struct `Category` dengan field `Id` dan `Name`.

### Helper
- **hello_world.go**: Fungsi sederhana `HelloWorld(name string) string` yang mengembalikan "Hello " + name.
- **hello_world_test.go**: Berisi unit test dan benchmark untuk fungsi HelloWorld, termasuk table-driven test dan penggunaan testify assert/require.

### Repository
- **category_repository.go**: Interface `CategoryRepository` dengan method `FindById(id string) *entity.Category`.
- **category_repository_mock.go**: Implementasi mock untuk `CategoryRepository` menggunakan testify/mock.

### Service
- **category_service.go**: Struct `CategoryService` yang mengimplementasikan business logic dengan dependency injection repository.
- **category_service_test.go**: Unit test untuk `CategoryService` menggunakan mock repository, termasuk test untuk success case dan not found case.

## Cara Menjalankan Test

### Menjalankan Semua Test
```bash
go test ./...
```

### Menjalankan Test dengan Verbose
```bash
go test -v ./...
```

### Menjalankan Test Spesifik
```bash
go test -v -run=TestCategoryService_GetNotFound
```

### Menjalankan Benchmark
```bash
go test -v -bench=.
```

### Menjalankan Benchmark Spesifik
```bash
go test -v -bench=BenchmarkTable
```

### Menjalankan Sub-Benchmark
```bash
go test -v -run=tidakAdaUnitTest -bench=BenchmarkTable/Eko
```

## Modul Go

- **Nama Modul**: `belajar-golang-unit-test`
- **Versi Go**: 1.15
- **Dependency**: `github.com/stretchr/testify v1.6.1`

## Persyaratan

- Go versi 1.15 atau lebih baru
- Dependency testify akan di-download otomatis saat `go mod tidy`

## Konsep yang Dipelajari

1. **Basic Unit Test**: Menggunakan `testing` package bawaan Go
2. **Testify Library**: Menggunakan `assert` dan `require` untuk assertion yang lebih baik
3. **Benchmark**: Mengukur performa fungsi dengan `testing.B`
4. **Table-Driven Test**: Menjalankan multiple test case dengan data table
5. **Mocking**: Menggunakan testify/mock untuk mock dependency
6. **Dependency Injection**: Memisahkan business logic dari dependency untuk memudahkan testing

## Catatan

File-file ini dirancang untuk pembelajaran bertahap tentang unit testing di Go. Mulai dari test sederhana hingga testing dengan mock untuk dependency external. Pastikan untuk menjalankan `go mod tidy` terlebih dahulu untuk menginstall dependency testify.