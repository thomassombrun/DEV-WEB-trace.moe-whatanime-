package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"trace.moe-whatanime/library"
)

func Favorites(w http.ResponseWriter, r *http.Request) {
	favorites, err := library.ListFavorites()
	if err != nil {
		http.Error(w, "Impossible de lire les favoris.", http.StatusInternalServerError)
		return
	}
	renderPage(w, "favorites.html", http.StatusOK, favorites)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 8192)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulaire invalide.", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.PostForm.Get("anilist_id"))
	title := strings.TrimSpace(r.PostForm.Get("title"))
	if err != nil || id <= 0 || title == "" {
		http.Error(w, "Anime invalide.", http.StatusBadRequest)
		return
	}
	if err := library.AddFavorite(library.Favorite{AniListID: id, Title: title}); err != nil {
		http.Error(w, "Impossible d'enregistrer le favori.", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 8192)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulaire invalide.", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.PostForm.Get("anilist_id"))
	if err != nil || id <= 0 {
		http.Error(w, "Anime invalide.", http.StatusBadRequest)
		return
	}
	if err := library.RemoveFavorite(id); err != nil {
		http.Error(w, "Impossible de retirer le favori.", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}
