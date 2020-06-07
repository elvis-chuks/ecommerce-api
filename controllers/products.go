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

		if len(product.Name) == 0{
			http.Error(w, `{"status":"error","msg":"product name can not be empty"}`,400)
			return
		}
		if len(product.Category) == 0{
			http.Error(w, `{"status":"error","msg":"product category can not be empty"}`,400)
			return
		}
		if len(product.Quantity) == 0{
			http.Error(w, `{"status":"error","msg":"product quantity can not be empty"}`,400)
			return
		}
		if len(product.Price) == 0{
			http.Error(w, `{"status":"error","msg":"product price can not be empty"}`,400)
			return
		}
		if len(product.Image) == 0{
			http.Error(w, `{"status":"error","msg":"product price can not be empty"}`,400)
			return
		}

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
		
		json.NewEncoder(w).Encode(models.Resp{"status":"success"})
		return
	}
	http.Error(w, `{"status":"error","msg":"method not allowed"}`,400)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	helpers.SetupResponse(&w,r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "GET"{
		db := helpers.InitDB()
		defer db.Close()

		query := fmt.Sprintf(`select id,category,name,quantity,price,image from products`)
		rows, err := db.Query(query)
		
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, `{"status":"error","msg":"bad request"}`,400)
			return
		}

		defer rows.Close()

		var products []models.Resp

		for rows.Next() {
			var id,category,name,quantity,price,image string
			rows.Scan(&id,&category,&name,&quantity,&price,&image)
			resp := models.Resp{
				"id":id,
				"category":category,
				"name":name,
				"quantity":quantity,
				"price":price,
				"image":image,
			}

			products = append(products,resp)
		}
		finalResponse := models.Resp{"status":"success","products":products}
		json.NewEncoder(w).Encode(finalResponse)
		return
	}
	http.Error(w, `{"status":"error","msg":"method not allowed"}`,400)
}

func GetProductsByCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	helpers.SetupResponse(&w,r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST"{
		var category models.Category
		_ = json.NewDecoder(r.Body).Decode(&category)

		db := helpers.InitDB()
		defer db.Close()

		query := fmt.Sprintf(`SELECT id,category,name,quantity,price,image from products where category = '%s'`,category.Id)

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, `{"status":"error","msg":"bad request"}`,400)
			return
		}

		defer rows.Close()

		var products []models.Resp

		for rows.Next() {
			var id,category,name,quantity,price,image string
			rows.Scan(&id,&category,&name,&quantity,&price,&image)
			resp := models.Resp{
				"id":id,
				"category":category,
				"name":name,
				"quantity":quantity,
				"price":price,
				"image":image,
			}

			products = append(products,resp)
		}

		finalResponse := models.Resp{"status":"success","products":products}
		json.NewEncoder(w).Encode(finalResponse)
		return
	}
	http.Error(w, `{"status":"error","msg":"method not allowed"}`,400)
	return
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	helpers.SetupResponse(&w,r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST"{
		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)

		if product.Id == 0{
			http.Error(w, `{"status":"error","msg":"product id can not be empty"}`,400)
			return
		}
		if len(product.Name) == 0{
			http.Error(w, `{"status":"error","msg":"product name can not be empty"}`,400)
			return
		}
		if len(product.Category) == 0{
			http.Error(w, `{"status":"error","msg":"product category can not be empty"}`,400)
			return
		}
		if len(product.Quantity) == 0{
			http.Error(w, `{"status":"error","msg":"product quantity can not be empty"}`,400)
			return
		}
		if len(product.Price) == 0{
			http.Error(w, `{"status":"error","msg":"product price can not be empty"}`,400)
			return
		}

		db := helpers.InitDB()
		defer db.Close()

		query := fmt.Sprintf(`UPDATE products set name='%s',category='%s',quantity='%s',price='%s' where id='%d'`,
		product.Name,product.Category,product.Quantity,product.Price,product.Id)

		_, err := db.Exec(query)
		
		if err != nil {
			http.Error(w, `{"status":"error","msg":"Something bad happened"}`,500)
			return
		}

		json.NewEncoder(w).Encode(models.Resp{"status":"success","details":models.Resp{"name":product.Name,"category":product.Category,"quantity":product.Quantity,"price":product.Price}})
		return
	}
	http.Error(w, `{"status":"error","msg":"method not allowed"}`,400)
	return
}