package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// sudah aman secara default, karena template sudah melakukan auto escape,
// sehingga kita tidak perlu khawatir dengan serangan XSS,
// karena template sudah melakukan escape terhadap karakter-karakter yang berbahaya, seperti <, >, &, ", ', dan lain-lain,
// sehingga ketika kita menampilkan data yang berasal dari user, maka data tersebut akan aman untuk ditampilkan di browser.
func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini Adalah Body<script>alert('Anda di Heck')</script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<h1>Ini Adalah Body</h1>"),
		//bisa juga css dan js yang berbahaya, seperti <script>alert('Anda di Heck')</script>, maka script tersebut akan di render sebagai HTML, sehingga script tersebut akan dieksekusi oleh browser, sehingga kita bisa terkena serangan XSS, karena kita bisa menyisipkan script berbahaya di dalam data yang kita kirim ke template.HTML.
		//jadi ketika kita menggunakan template.HTML, maka data yang kita kirim ke template.HTML akan di render sebagai HTML, sehingga kita harus berhati-hati ketika menggunakan template.HTML, karena jika kita tidak hati-hati, maka kita bisa terkena serangan XSS, karena kita bisa menyisipkan script berbahaya di dalam data yang kita kirim ke template.HTML.
		//jadi ketika kita menampilkan data yang berasal dari user, maka data tersebut akan di render sebagai HTML, sehingga kita harus berhati-hati ketika menggunakan template.HTML, karena jika kita tidak hati-hati, maka kita bisa terkena serangan XSS, karena kita bisa menyisipkan script berbahaya di dalam data yang kita kirim ke template.HTML.
		//ini akan di render sebagai HTML, sehingga tag <h1> akan di render sebagai heading 1, bukan sebagai string biasa,
		// sehingga ketika kita menampilkan data yang berasal dari user, maka data tersebut akan di render sebagai HTML,
		// sehingga kita harus berhati-hati ketika menggunakan template.HTML, karena jika kita tidak hati-hati, maka kita bisa terkena serangan XSS, karena kita bisa menyisipkan script berbahaya di dalam data yang kita kirim ke template.HTML.
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
