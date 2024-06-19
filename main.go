package main

import (
	"cmp"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed index.gohtml
var Assets embed.FS

func main() {
	mux := http.NewServeMux()

	tmpl := template.Must(template.ParseFS(Assets, "index.gohtml"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, nil)
	})

	port := cmp.Or(os.Getenv("PORT"), "8080")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalln(fmt.Errorf("error on listen and serve http: %w", err))
	}
}
