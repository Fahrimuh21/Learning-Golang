package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko')" //insert, update, delete itu termasuk DML (Data Manipulation Language) yang
	// digunakan untuk memanipulasi data di database jadi disini bisa kita gunakan context untuk mengatur batas waktu eksekusi, misal kita kasih timeout 5
	// detik, maka jika proses eksekusi lebih dari 5 detik maka proses akan dibatalkan (cancel) dan error context deadline exceeded akan muncul
	// sedangkan create table, alter table, drop table itu termasuk DDL (Data Definition Language) yang digunakan untuk mendefinisikan struktur database
	//_, err := db.Exec(script) tanpa context, maka proses eksekusi akan terus berjalan sampai selesai
	//tidak menggunakan context, maka proses eksekusi akan terus berjalan sampai selesai
	//walaupun sudah tidak dibutuhkan lagi, misal user sudah menutup aplikasi atau sudah pindah ke halaman lain
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	//rows kaya cursor yang digunakan untuk menampung hasil query,
	//jadi kita bisa menggunakan rows untuk mengambil data dari database dengan cara iterasi menggunakan rows.Next() dan mengambil data dengan rows.Scan()
	if err != nil {
		panic(err)
	}
	defer rows.Close() //untuk memastikan resource yang digunakan untuk menampung hasil query dibebaskan
	//setelah selesai digunakan, jika tidak maka akan terjadi memory leak karena resource akan terus digunakan sampai program selesai dijalankan

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		//simpan id dan name ke variable id dan name, jika terjadi error maka akan muncul error, misal jika tipe data di database tidak sesuai dengan tipe data di variable maka akan muncul error
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString //menangani data yang nullable, jika data di database null maka variable email akan bernilai null, jika data di database tidak null maka variable email akan bernilai string
		var balance int32
		var rating float64
		var birthDate sql.NullTime //menangani data yang nullable, jika data di database null maka variable birthDate akan bernilai null, jika data di database tidak null maka variable birthDate akan bernilai time.Time
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		//simpan data ke variable yang sudah didefinisikan, jika terjadi error maka akan muncul error, misal jika tipe data di database tidak sesuai dengan tipe data di variable maka akan muncul error
		if err != nil {
			panic(err)
		}
		fmt.Println("================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid { //untuk mengecek apakah data email valid atau tidak, jika valid maka data di database tidak null, jika tidak valid maka data di database null
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid { //untuk mengecek apakah data birthDate valid atau tidak, jika valid maka data di database tidak null, jika tidak valid maka data di database null
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	//jadi setiap query yang menggunakan parameter, maka parameter tersebut akan di escape terlebih dahulu sebelum digunakan dalam query,
	// sehingga jika ada karakter khusus seperti ' maka akan di escape menjadi \' sehingga tidak akan mempengaruhi query yang dijalankan,
	// dengan menggunakan parameter maka kita bisa mencegah terjadinya sql injection karena parameter akan di escape terlebih dahulu sebelum digunakan dalam query
	//? itu adalah placeholder untuk parameter yang akan diisi dengan nilai username dan password, jadi ketika query dijalankan maka parameter tersebut akan diisi dengan nilai username dan password yang sudah didefinisikan sebelumnya
	//bisa pakai insert into user(username, password) values(?, ?) untuk insert data dengan parameter, jadi setiap parameter yang digunakan dalam query akan di escape terlebih dahulu sebelum digunakan dalam query sehingga bisa mencegah terjadinya sql injection
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "eko'; DROP TABLE user; #"
	password := "eko"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "eko@gmail.com"
	comment := "Test komen"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId() //id terakhir yang berhasil di insert, jadi jika kita insert data baru maka id yang dihasilkan akan lebih besar dari id sebelumnya, dengan menggunakan auto increment maka kita tidak perlu lagi mengatur id secara manual karena id akan di generate secara otomatis oleh database
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		//jadi dengan menggunakan prepared statement kita bisa mengeksekusi query yang sama dengan parameter yang berbeda tanpa harus membuat query baru setiap kali ingin mengeksekusi query, sehingga bisa meningkatkan performa karena query yang sudah di prepare akan di compile terlebih dahulu oleh database sehingga ketika kita eksekusi query dengan parameter yang berbeda maka database tidak perlu lagi mengcompile query tersebut karena sudah di compile sebelumnya,
		// jadi kita hanya perlu mengisi parameter yang berbeda setiap kali ingin mengeksekusi query, dengan menggunakan prepared statement kita juga bisa mencegah terjadinya sql injection karena parameter yang digunakan dalam query akan di escape terlebih dahulu sebelum digunakan dalam query sehingga tidak akan mempengaruhi query yang dijalankan
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// do transaction
	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}

	err = tx.Commit()
	//gagalin = rollback, jadi data yang sudah di insert tidak akan masuk ke database, dengan menggunakan transaction kita bisa memastikan bahwa semua proses yang dilakukan dalam transaction berhasil atau tidak sama sekali, jadi jika ada salah satu proses yang gagal maka semua proses yang sudah dilakukan dalam transaction akan dibatalkan (rollback) sehingga data di database tetap konsisten
	//berhasil = commit, jadi data yang sudah di insert akan masuk ke database, dengan menggunakan transaction kita bisa memastikan bahwa semua proses yang dilakukan dalam transaction berhasil atau tidak sama sekali, jadi jika semua proses berhasil maka semua proses yang sudah dilakukan dalam transaction akan disimpan (commit) ke database sehingga data di database tetap konsisten
	if err != nil {
		panic(err)
	}
}
