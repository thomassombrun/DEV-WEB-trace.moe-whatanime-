package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"trace.moe-whatanime/api"
	"trace.moe-whatanime/models"
)

type HomeData struct {
	Quota models.Quota
	Error string
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := HomeData{}

	quota, err := api.GetQuota()
	if err != nil {
		fmt.Println("Erreur API :", err)
		data.Error = "Impossible de récupérer les informations de l'API pour le moment."
	} else {
		data.Quota = quota
	}
	tmpl, err := template.ParseFiles(
		"templates/index.html",
		"templates/partials/header.html",
		"templates/partials/footer.html",
	)
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Erreur template:", err)
		http.Error(w, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
		return
	}
}
