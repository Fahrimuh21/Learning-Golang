package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout) //membuat csv writer yang menulis ke os.Stdout (console)

	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"}) //menulis data csv
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush() //memastikan semua data tertulis
}
