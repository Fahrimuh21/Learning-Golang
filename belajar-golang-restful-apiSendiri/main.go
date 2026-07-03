package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

// @title Category API
// @version 1.0
// @description API for categories
// func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("OK"))
// }

// func main() {
// 	router := http.NewServeMux()
// 	router.HandleFunc("/api/categories", getCategoriesHandler)
// 	http.ListenAndServe(":3000", router)
// }

func main() {
	db := app.NewDB()
	defer db.Close()
	validator := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:CategoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:CategoryId", categoryController.Update)
	router.DELETE("/api/categories/:CategoryId", categoryController.Delete)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
