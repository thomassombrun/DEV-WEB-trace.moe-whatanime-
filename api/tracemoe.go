package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"trace.moe-whatanime/models"
)

const baseURL = "https://api.trace.moe"

var client = &http.Client{Timeout: 10 * time.Second}

func GetQuota() (models.Quota, error) {
	var quota models.Quota

	resp, err := client.Get(baseURL + "/me")
	if err != nil {
		return quota, fmt.Errorf("API injoignable : %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return quota, fmt.Errorf("l'API a répondu avec le code %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&quota)
	if err != nil {
		return quota, fmt.Errorf("réponse JSON invalide : %w", err)
	}

	return quota, nil
}

func SearchImage(ctx context.Context, image []byte) (*models.SceneMatch, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/search?anilistInfo", bytes.NewReader(image))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API injoignable : %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("l'API a répondu avec le code %d", resp.StatusCode)
	}

	var data struct {
		Error  string `json:"error"`
		Result []struct {
			AniList    json.RawMessage `json:"anilist"`
			Filename   string          `json:"filename"`
			Episode    any             `json:"episode"`
			From       float64         `json:"from"`
			To         float64         `json:"to"`
			Similarity float64         `json:"similarity"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("réponse JSON invalide : %w", err)
	}
	if data.Error != "" {
		return nil, fmt.Errorf("erreur trace.moe : %s", data.Error)
	}
	if len(data.Result) == 0 {
		return nil, nil
	}

	best := data.Result[0]
	for _, result := range data.Result[1:] {
		if result.Similarity > best.Similarity {
			best = result
		}
	}

	title := best.Filename
	var anime struct {
		Title struct {
			English string `json:"english"`
			Romaji  string `json:"romaji"`
			Native  string `json:"native"`
		} `json:"title"`
	}
	// L'API peut renvoyer seulement l'identifiant AniList pour un résultat.
	if err := json.Unmarshal(best.AniList, &anime); err == nil {
		switch {
		case anime.Title.English != "":
			title = anime.Title.English
		case anime.Title.Romaji != "":
			title = anime.Title.Romaji
		case anime.Title.Native != "":
			title = anime.Title.Native
		}
	}

	episode := ""
	if best.Episode != nil {
		episode = fmt.Sprint(best.Episode)
	}
	return &models.SceneMatch{
		Title:             title,
		Episode:           episode,
		From:              best.From,
		To:                best.To,
		SimilarityPercent: best.Similarity * 100,
	}, nil
}
