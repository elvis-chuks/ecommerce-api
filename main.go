package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elvis-chuks/ecommerce-api/controllers"
	"github.com/elvis-chuks/ecommerce-api/helpers"
)

func test(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "works", 200)
}

func main() {
	er := helpers.InitTables()

	if er != nil {
		panic(er)
	}

	http.HandleFunc("/", test)
	http.HandleFunc("/v1/addproduct", controllers.AddProduct)
	http.HandleFunc("/v1/getallproducts", controllers.GetAllProducts)
	http.HandleFunc("/v1/getproductsbycategory", controllers.GetProductsByCategory)
	http.HandleFunc("/v1/updateproduct", controllers.UpdateProduct)
	http.HandleFunc("/v1/deleteproduct", controllers.DeleteProduct)
	http.HandleFunc("/v1/addcategory", controllers.AddCategory)
	http.HandleFunc("/v1/getcategories", controllers.GetAllCategories)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	port = fmt.Sprintf(":"+"%s", port)

	fmt.Println(fmt.Sprintf("app running on http://localhost%s", port))

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
	}
}
