package handlers

import "net/http"

func Search(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "search.html", http.StatusOK, nil)
}
