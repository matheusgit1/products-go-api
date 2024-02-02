package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"https://github.com/matheusgit1/products-go-api/internal/database"
	"https://github.com/matheusgit1/products-go-api/internal/service"
	"https://github.com/matheusgit1/products-go-api/internal/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/commerce")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
