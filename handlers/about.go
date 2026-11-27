package handlers

import "net/http"

func About(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "about.html", http.StatusOK, nil)
}
