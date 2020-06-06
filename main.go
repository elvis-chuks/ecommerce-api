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
	

	http.HandleFunc("/", test)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	port = fmt.Sprintf(":"+"%s", port)

	fmt.Println(fmt.Sprintf("app running on http://localhost%s",port))

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
	}
}
