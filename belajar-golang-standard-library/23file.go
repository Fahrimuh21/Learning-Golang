package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil { // jika terjadi error saat membuka file yaitu ketika (file sudah ada => eror = true)
		return err
	}
	defer file.Close()        // pastikan file ditutup setelah selesai digunakan
	file.WriteString(message) // tulis message ke file
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file) //membuat buffered reader dari file
	var message string              //menyimpan hasil bacaan
	for {
		line, _, err := reader.ReadLine() //baca per line
		if err == io.EOF {
			break
		}
		message += string(line) + "\n" //gabungkan hasil baca ke message
	}
	return message, nil //kembalikan message dan nil error
}

func main() {
	//createNewFile("sample.log", "this is sample log")

	//result, _ := readFile("sample.log")
	//fmt.Println(result)

	addToFile("sample.log", "\nthis is add message")
}
