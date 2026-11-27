package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"trace.moe-whatanime/library"
	"trace.moe-whatanime/models"
)

type SearchData struct {
	Result      *models.SceneMatch
	Message     string
	ImageType   string
	ImageBase64 string
}

func Search(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodHead:
		renderPage(w, "search.html", http.StatusOK, SearchData{})
	case http.MethodPost:
		searchImage(w, r)
	default:
		w.Header().Set("Allow", "GET, HEAD, POST")
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func searchImage(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	file, _, err := r.FormFile("image")
	if err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			renderPage(w, "search.html", http.StatusRequestEntityTooLarge, SearchData{Message: "L'envoi dépasse la limite de 10 Mo."})
			return
		}
		renderPage(w, "search.html", http.StatusBadRequest, SearchData{Message: "Choisis une image à envoyer."})
		return
	}
	defer file.Close()

	image, err := io.ReadAll(file)
	imageType := http.DetectContentType(image)
	if err != nil || !strings.HasPrefix(imageType, "image/") {
		renderPage(w, "search.html", http.StatusBadRequest, SearchData{Message: "Le fichier doit être une image valide."})
		return
	}

	result, err := library.SearchImage(r.Context(), image)
	if err != nil {
		fmt.Println("Erreur recherche :", err)
		renderPage(w, "search.html", http.StatusBadGateway, SearchData{Message: "La recherche est indisponible pour le moment."})
		return
	}
	if result == nil {
		renderPage(w, "search.html", http.StatusOK, SearchData{Message: "Aucun résultat trouvé."})
		return
	}

	renderPage(w, "search.html", http.StatusOK, SearchData{
		Result:      result,
		ImageType:   imageType,
		ImageBase64: base64.StdEncoding.EncodeToString(image),
	})
}
