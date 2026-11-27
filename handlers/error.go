package handlers

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "error.html", http.StatusNotFound, nil)
}
