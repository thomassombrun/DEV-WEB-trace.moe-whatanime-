package handlers

import "net/http"

func Collection(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "collection.html", http.StatusOK, nil)
}
