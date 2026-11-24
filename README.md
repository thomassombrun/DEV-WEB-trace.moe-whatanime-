# DEV-WEB-trace.moe-whatanime-

## Routes

| Méthode | Route | Handler | Template | Description |
|---|---|---|---|---|
| GET | `/` | `home.go` | `index.html` | Page d'accueil : présentation de l'univers et du thème du site |
| GET | `/collection` | `collection.go` | `collection.html` | Liste des animes avec filtres cumulables et pagination (`?page=1&limit=10&...`) |
| GET | `/anime/{id}` | `detail.go` | `detail.html` | Détails d'un anime à partir de son identifiant |
| GET | `/search` | `search.go` | `search.html` | Résultats d'une recherche (`?q=...`) sur au moins deux propriétés |
| GET | `/favorites` | `favorites.go` | `favorites.html` | Affiche la liste des favoris |
| POST | `/favorites/add` | `favorites.go` | — | Ajoute un anime aux favoris (`data/favorites.json`), puis redirige |
| POST | `/favorites/delete` | `favorites.go` | — | Supprime un anime des favoris, puis redirige |
| GET | `/about` | `about.go` | `about.html` | Page à-propos : retour d'expérience sous forme de FAQ |
| GET | `/static/...` | — | — | Fichiers statiques (CSS, images) |
| — | toute autre route | `error.go` | `error.html` | Page d'erreur 404 / erreurs serveur et API |