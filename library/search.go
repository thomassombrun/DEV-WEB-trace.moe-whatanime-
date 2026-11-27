package library

import (
	"context"
	"encoding/json"
	"fmt"

	"trace.moe-whatanime/api"
	"trace.moe-whatanime/models"
)

func SearchImage(ctx context.Context, image []byte) (*models.SceneMatch, error) {
	results, err := api.SearchScenes(ctx, image)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}

	best := results[0]
	for _, result := range results[1:] {
		if result.Similarity > best.Similarity {
			best = result
		}
	}

	title := best.Filename
	var anime struct {
		ID    int `json:"id"`
		Title struct {
			English string `json:"english"`
			Romaji  string `json:"romaji"`
			Native  string `json:"native"`
		} `json:"title"`
	}
	if err := json.Unmarshal(best.AniList, &anime); err == nil {
		switch {
		case anime.Title.English != "":
			title = anime.Title.English
		case anime.Title.Romaji != "":
			title = anime.Title.Romaji
		case anime.Title.Native != "":
			title = anime.Title.Native
		}
	} else {
		json.Unmarshal(best.AniList, &anime.ID)
	}

	episode := ""
	if best.Episode != nil {
		episode = fmt.Sprint(best.Episode)
	}
	return &models.SceneMatch{
		AniListID:         anime.ID,
		Title:             title,
		Episode:           episode,
		From:              best.From,
		To:                best.To,
		SimilarityPercent: best.Similarity * 100,
		ImageURL:          best.Image,
	}, nil
}
