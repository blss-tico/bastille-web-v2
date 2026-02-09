package web

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func renderTemplateUtil(w http.ResponseWriter, name string, data any) {
	log.Println("renderTemplate")

	tmplt := "./web/templates/" + name
	files := []string{
		"./web/templates/base.html",
		"./web/templates/partials/navbar.html",
		"./web/templates/partials/sidebar.html",
	}
	files = append(files, tmplt)

	var err error
	templates, err = template.ParseFiles(files...)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	err = templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
