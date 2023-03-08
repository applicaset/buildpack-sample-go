package main

import (
	"embed"
	"fmt"
	"github.com/nasermirzaei89/env"
	"html/template"
	"log"
	"net/http"
)

//go:embed index.gohtml
var Assets embed.FS

func main() {
	mux := http.NewServeMux()

	tmpl := template.Must(template.ParseFiles("index.gohtml"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, nil)
	})

	port := env.GetString("PORT", "8080")

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalln(fmt.Errorf("error on listen and serve http: %w", err))
	}
}
