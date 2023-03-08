package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/nasermirzaei89/env"
)

//go:embed index.gohtml
var Assets embed.FS

func main() {
	mux := http.NewServeMux()

	tmpl := template.Must(template.ParseFS(Assets, "index.gohtml"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, nil)
	})

	port := env.GetString("PORT", "8080")

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalln(fmt.Errorf("error on listen and serve http: %w", err))
	}
}
