package main

import (
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/app"
	"programmerzamannow/belajar-golang-restful-api/controller"
	"programmerzamannow/belajar-golang-restful-api/exception"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/middleware"
	"programmerzamannow/belajar-golang-restful-api/repository"
	"programmerzamannow/belajar-golang-restful-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.FindById)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
