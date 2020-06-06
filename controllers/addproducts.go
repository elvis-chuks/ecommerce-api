package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/elvis-chuks/ecommerce-api/models"
	"github.com/elvis-chuks/ecommerce-api/helpers"
)


func AddProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	helpers.SetupResponse(&w,r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST"{
		var product models.Product

		_ = json.NewDecoder(r.Body).Decode(&product)

		link, err := helpers.Upload(product.Image)
		
		if err != nil {
			http.Error(w,fmt.Sprintf(`{"status":"error","msg":%s}`,err.Error()),400)
		}

		// enter product into products table
		query := fmt.Sprintf(`INSERT INTO PRODUCTS(name,category,quantity,price,image) VALUES ('%s','%s','%s','%s','%s')`,
			product.Name,product.Category,product.Quantity,product.Price,link)
		
		db := helpers.InitDB()
		defer db.Close()

		_, err = db.Exec(query)

		if err != nil {
			fmt.Println(err)
			http.Error(w,fmt.Sprintf(`{"status":"error","msg":%s}`,err.Error()),400)
		}

		json.NewEncoder(w).Encode(`{"status":"success"}`)
		return
	}
	http.Error(w, `{"status":"error","msg":"method not allowed"}`,400)
}