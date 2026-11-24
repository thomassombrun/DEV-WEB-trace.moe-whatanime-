package api

import (
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
