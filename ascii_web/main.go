package main

import (
	"fmt"
	//"html/template"
	"net/http"
)

func ascii_art_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
            return
		}
		text := r.FormValue("text")
		banner := r.FormValue("banner")
	}
}

func main() {
	fmt.Println("Server Online")
	http.ListenAndServe(":8080", nil)
}

