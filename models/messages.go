package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromID    uint   `json:"from_id"`
	ToID      uint   `json:"to_id"`
	Type      string   `json:"type"`
	Media     int   `json:"media"`
	Content   string   `json:"content"`
	Picture   string   `json:"picture"`
	Url       string   `json:"url"`
	Description string `json:"description"`
	Amount    int    `json:"amount"`
}

func (m *Message) TableName() string {
	return "messages"
}
