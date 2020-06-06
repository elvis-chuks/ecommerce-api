package main

import (
	"fmt"
	"net/http"
	"os"
)

func test(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "works", 200)
}

func main() {
	fmt.Println("Ecommerce api")

	http.HandleFunc("/", test)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	port = fmt.Sprintf(":"+"%s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
	}
}
