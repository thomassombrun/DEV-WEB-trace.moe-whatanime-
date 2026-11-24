package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/index.html",
		"templates/partials/header.html",
		"templates/partials/footer.html",
	)
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Erreur template:", err)
		http.Error(w, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
		return
	}
}
