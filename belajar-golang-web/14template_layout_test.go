package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Eko",
	})
}

// tentang penggabungan template, kita bisa menggunakan template yang sudah kita buat sebelumnya, seperti header dan footer, lalu kita gabungkan dengan template layout, sehingga kita bisa membuat tampilan yang lebih kompleks dengan mudah.
//gabungin file html yang sudah dibuat sebelumnya dengan template layout, lalu kita panggil template layout dengan nama "layout", lalu kita kirim data ke template layout dengan menggunakan map[string]interface{} yang berisi data yang ingin kita kirim ke template layout, seperti Title dan Name.

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
