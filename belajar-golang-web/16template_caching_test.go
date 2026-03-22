package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

//kita embed dulu biar lebih mudah untuk mengelola template kita, dengan menggunakan embed
// kita bisa menyimpan semua template kita di dalam folder templates, lalu kita bisa memanggil template yang kita inginkan dengan menggunakan nama file template tersebut, seperti "simple.gohtml" pada contoh di bawah ini.

// mantappppuuuu jiwaaa
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

// semua template diatas tidak efektif karna parsing terus idealnya parsing diawal
func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
