package handlers

import (
	"fmt"
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
	renderPage(w, "index.html", http.StatusOK, data)
}
