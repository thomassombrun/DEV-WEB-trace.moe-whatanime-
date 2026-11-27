package library

import (
	"bytes"
	"encoding/json"
	"os"
	"sync"
)

const favoritesFile = "data/favorites.json"

var favoritesMutex sync.Mutex

type Favorite struct {
	AniListID int    `json:"anilist_id"`
	Title     string `json:"title"`
}

func ListFavorites() ([]Favorite, error) {
	favoritesMutex.Lock()
	defer favoritesMutex.Unlock()

	return readFavorites()
}

func readFavorites() ([]Favorite, error) {
	data, err := os.ReadFile(favoritesFile)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if len(bytes.TrimSpace(data)) == 0 {
		return nil, nil
	}
	var favorites []Favorite
	err = json.Unmarshal(data, &favorites)
	return favorites, err
}

func AddFavorite(anime Favorite) error {
	favoritesMutex.Lock()
	defer favoritesMutex.Unlock()

	favorites, err := readFavorites()
	if err != nil {
		return err
	}
	for _, saved := range favorites {
		if saved.AniListID == anime.AniListID {
			return nil
		}
	}

	return writeFavorites(append(favorites, anime))
}

func RemoveFavorite(id int) error {
	favoritesMutex.Lock()
	defer favoritesMutex.Unlock()

	favorites, err := readFavorites()
	if err != nil {
		return err
	}
	for index, anime := range favorites {
		if anime.AniListID == id {
			favorites = append(favorites[:index], favorites[index+1:]...)
			return writeFavorites(favorites)
		}
	}
	return nil
}

func writeFavorites(favorites []Favorite) error {
	data, err := json.MarshalIndent(favorites, "", "  ")
	if err != nil {
		return err
	}
	tempFile := favoritesFile + ".tmp"
	defer os.Remove(tempFile)
	if err := os.WriteFile(tempFile, append(data, '\n'), 0644); err != nil {
		return err
	}
	return os.Rename(tempFile, favoritesFile)
}
