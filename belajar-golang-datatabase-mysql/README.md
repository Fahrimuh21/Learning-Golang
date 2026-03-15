# Belajar Golang Database MySQL

Repositori ini berisi materi pembelajaran tentang penggunaan database MySQL dalam bahasa pemrograman Go. Folder ini mencakup koneksi database, eksekusi SQL, dan implementasi repository pattern.

## Struktur Folder

```
belajar-golang-database/
├── go.mod                          # File modul Go dengan dependency MySQL driver
├── .gitignore                      # File gitignore
├── database.go                     # Setup koneksi database MySQL
├── database_test.go                # Test koneksi database
├── sql_test.go                     # Test eksekusi SQL (Exec, Query)
├── entity/                         # Package entity
│   └── comment.go                  # Struct Comment
└── repository/                     # Package repository
    ├── comment_repository.go       # Interface CommentRepository
    ├── comment_repository_impl.go  # Implementasi CommentRepository
    └── comment_repository_impl_test.go # Test untuk repository implementation
```

## Deskripsi File

### Root Files
- **go.mod**: Mendefinisikan modul `belajar-golang-database` dengan dependency `github.com/go-sql-driver/mysql v1.5.0`.
- **database.go**: Fungsi `GetConnection()` untuk setup koneksi MySQL dengan connection pooling.
- **database_test.go**: Test untuk membuka dan menutup koneksi database.
- **sql_test.go**: Test untuk eksekusi SQL menggunakan `ExecContext` dan `QueryContext` dengan context.

### Entity
- **comment.go**: Struct `Comment` dengan field `Id`, `Email`, dan `Comment`.

### Repository
- **comment_repository.go**: Interface `CommentRepository` dengan method `Insert`, `FindById`, dan `FindAll`.
- **comment_repository_impl.go**: Implementasi concrete dari `CommentRepository` dengan operasi CRUD ke database.
- **comment_repository_impl_test.go**: Unit test untuk testing repository operations.

## Setup Database

Sebelum menjalankan test, pastikan:

1. MySQL server sudah terinstall dan running
2. Database `belajar_golang_database` sudah dibuat
3. Tabel `comments` sudah dibuat dengan struktur:
   ```sql
   CREATE TABLE comments (
       id INT AUTO_INCREMENT PRIMARY KEY,
       email VARCHAR(255) NOT NULL,
       comment TEXT NOT NULL
   );
   ```

## Cara Menjalankan Test

### Menjalankan Semua Test
```bash
go test -v ./...
```

### Menjalankan Test Spesifik
```bash
go test -v -run=TestCommentInsert
```

### Menjalankan Test dengan Timeout
```bash
go test -v -timeout=30s
```

## Modul Go

- **Nama Modul**: `belajar-golang-database`
- **Versi Go**: 1.15
- **Dependency**: `github.com/go-sql-driver/mysql v1.5.0`

## Persyaratan

- Go versi 1.15 atau lebih baru
- MySQL server
- Database `belajar_golang_database` dengan tabel `comments`

## Konsep yang Dipelajari

1. **Database Connection**: Setup koneksi MySQL dengan connection pooling
2. **SQL Execution**: Menggunakan `ExecContext` untuk DML dan `QueryContext` untuk queries
3. **Context Usage**: Menggunakan context untuk timeout dan cancellation
4. **Repository Pattern**: Abstraction layer antara business logic dan data access
5. **CRUD Operations**: Create, Read, Update, Delete operations
6. **Error Handling**: Proper error handling untuk database operations
7. **Resource Management**: Menggunakan `defer` untuk cleanup resources

## Catatan

Folder ini mendemonstrasikan best practices untuk database programming di Go:
- Menggunakan context untuk timeout dan cancellation
- Connection pooling untuk performance
- Repository pattern untuk separation of concerns
- Proper error handling dan resource cleanup

Pastikan MySQL server sudah dikonfigurasi dengan benar sebelum menjalankan test. Beberapa test mungkin mengubah data di database, jadi gunakan database development atau test database.