package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func renderPage(w http.ResponseWriter, filename string, status int, data any) {
	tmpl, err := template.ParseFiles(
		"templates/"+filename,
		"templates/partials/header.html",
		"templates/partials/footer.html",
	)
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	var page bytes.Buffer
	err = tmpl.Execute(&page, data)
	if err != nil {
		fmt.Println("Erreur template:", err)
		http.Error(w, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	w.Write(page.Bytes())
}
