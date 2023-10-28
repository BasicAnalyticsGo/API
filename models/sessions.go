package models

type Session struct {
	ID        uint   `gorm:"primaryKey"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"` // Should be encoded
	Country   string `json:"country"`
	City      string `json:"city"`
}
