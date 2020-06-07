package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/elvis-chuks/ecommerce-api/models"
	"github.com/elvis-chuks/ecommerce-api/helpers"
)

func AddCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	helpers.SetupResponse(&w,r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST"{
		var category models.Category
		_ = json.NewDecoder(r.Body).Decode(&category)

		if len(category.Name) != 0 {
			db := helpers.InitDB()
			defer db.Close()

			query := fmt.Sprintf(`INSERT INTO categories(name) VALUES('%s')`,category.Name)
			_, err := db.Exec(query)

			if err != nil {
				panic(err)
				http.Error(w, `{"status":"error","msg":"Something bad happened"}`,500)
				return
			}
			
			json.NewEncoder(w).Encode(models.Resp{"status":"success"})
			return
		}
		http.Error(w, `{"status":"error","msg":"bad request"}`,400)
		return
	}
	http.Error(w, `{"status":"error","msg":"method not allowed"}`,400)
	return
}
func GetAllCategories(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	helpers.SetupResponse(&w,r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "GET"{
		db := helpers.InitDB()
		defer db.Close()

		query := fmt.Sprintf(`select id,name from categories`)
		rows, err := db.Query(query)

		if err != nil {
			http.Error(w, `{"status":"error","msg":"Something bad happened"}`,500)
			return
		}
		
		defer rows.Close()

		var categories []models.Resp

		for rows.Next() {
			var id, name string
			rows.Scan(&id,&name)
			categories = append(categories,models.Resp{"id":id,"name":name})
		}
		json.NewEncoder(w).Encode(models.Resp{"status":"success","categories":categories})
		return
	}
	http.Error(w, `{"status":"error","msg":"method not allowed"}`,400)
	return
}
