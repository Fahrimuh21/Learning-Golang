package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatern(t *testing.T) {
	router := httprouter.New()

	//params itu buat kaya indikator gitu
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		ItemId := params.ByName("itemId")
		text := "Product " + id + " Items " + ItemId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Items 1", string(body))
}

func TestRouterPaternCatchAll(t *testing.T) {
	router := httprouter.New()

	//params itu buat kaya indikator gitu
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		imageId := params.ByName("image")
		text := "image : " + imageId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profiles.jpg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "image : /small/profiles.jpg", string(body))
}
