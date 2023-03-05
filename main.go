package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Welcome")

	})

	fmt.Println("listening at 8080")
	http.ListenAndServe(":8080", nil)
}
