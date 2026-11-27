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

type SearchResult struct {
	AniList    json.RawMessage `json:"anilist"`
	Filename   string          `json:"filename"`
	Episode    any             `json:"episode"`
	From       float64         `json:"from"`
	To         float64         `json:"to"`
	Similarity float64         `json:"similarity"`
	Image      string          `json:"image"`
}

func SearchScenes(ctx context.Context, image []byte) ([]SearchResult, error) {
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
		Error  string         `json:"error"`
		Result []SearchResult `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("réponse JSON invalide : %w", err)
	}
	if data.Error != "" {
		return nil, fmt.Errorf("erreur trace.moe : %s", data.Error)
	}
	return data.Result, nil
}
