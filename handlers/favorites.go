package handlers

import "net/http"

func Favorites(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "favorites.html", http.StatusOK, nil)
}
