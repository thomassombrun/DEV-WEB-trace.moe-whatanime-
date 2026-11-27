package handlers

import "net/http"

func Detail(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "detail.html", http.StatusOK, nil)
}
