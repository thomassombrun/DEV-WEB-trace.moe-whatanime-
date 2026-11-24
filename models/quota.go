package models

type Quota struct {
	ID          string `json:"id"`
	Priority    int    `json:"priority"`
	Concurrency int    `json:"concurrency"`
	Quota       int    `json:"quota"`
	QuotaUsed   int    `json:"quota_used"`
}
