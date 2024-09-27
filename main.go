package main

import (
	"fmt"
	"html/template"
	"net/http"

	fs "fs/ascii"
)

type PageVariables struct {
	Input  string
	Result string
}

var (
	vars PageVariables
	tpl  *template.Template
)

func main() {
	var err error
	tpl, err = template.ParseGlob("html/*.html")
	if err != nil {
		panic(err)
	}
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/art", ProcessForm)
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	vars = PageVariables{
		Input:  "",
		Result: "",
	}
	RenderTemplate(w, vars)
}

func ProcessForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		banner := r.FormValue("banner")
		input := r.FormValue("text")

		result := fs.FinalPrint(input, banner)

		vars = PageVariables{
			Input:  input,
			Result: result,
		}
		RenderTemplate(w, vars)
	}
}

func RenderTemplate(w http.ResponseWriter, vars PageVariables) {
	if err := tpl.ExecuteTemplate(w, "index.html", vars); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}
}
