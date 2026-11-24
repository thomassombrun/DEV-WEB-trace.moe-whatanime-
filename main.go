package main

import (
	"fmt"
	"net/http"

	"trace.moe-whatanime/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	fmt.Println("Le serveur est allumé sur http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}
