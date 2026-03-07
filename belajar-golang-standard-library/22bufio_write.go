package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)       //membuat buffered writer yang menulis ke os.Stdout (console)
	_, _ = writer.WriteString("hello world\n") //menulis string ke buffer
	_, _ = writer.WriteString("Selamat belajar\n")
	writer.Flush() //memastikan semua data tertulis ke output
}
