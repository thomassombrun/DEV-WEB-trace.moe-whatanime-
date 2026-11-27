package main

import (
	"fmt"
	"net/http"

	"trace.moe-whatanime/handlers"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/collection", handlers.Collection)
	mux.HandleFunc("/anime/{id}", handlers.Detail)
	mux.HandleFunc("/favorites", handlers.Favorites)
	mux.HandleFunc("POST /favorites/add", handlers.AddFavorite)
	mux.HandleFunc("POST /favorites/delete", handlers.RemoveFavorite)
	mux.HandleFunc("/search", handlers.Search)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", handlers.NotFound)
	return mux
}

func main() {
	fmt.Println("Le serveur est allumé sur http://localhost:8080")

	err := http.ListenAndServe(":8080", routes())
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}
