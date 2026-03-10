package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//t, err := template.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}

	t := template.Must(template.New("SIMPLE").Parse(templateText))
	//template.Must digunakan untuk membuat template dan langsung
	// mem-parsing template text yang sudah kita buat sebelumnya.
	// Jika terjadi error saat parsing, maka template.Must akan panic
	// dan menghentikan program. Dengan cara ini, kita bisa memastikan
	// bahwa template yang kita buat sudah benar dan tidak terjadi error saat parsing.

	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

// satu direktori bisa berisi banyak file template, kita bisa menggunakan template.ParseGlob untuk mem-parsing semua file template yang ada di dalam direktori tersebut. Dalam contoh ini, kita mem-parsing semua file template yang ada di dalam direktori templates dengan menggunakan pattern *.gohtml. Dengan cara ini, kita bisa mengelompokkan file-file template yang kita buat ke dalam satu direktori dan mem-parsing semuanya sekaligus dengan menggunakan template.ParseGlob.
func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
