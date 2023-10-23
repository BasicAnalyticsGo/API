package models

type Session struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"IP"` // Should be encoded
}
