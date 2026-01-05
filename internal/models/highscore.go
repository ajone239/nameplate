package models

import "time"

type HighScore struct {
	Id         int       `json:"id"`
	PlayerName string    `json:"player_name"`
	GameName   string    `json:"game_name"`
	Score      int       `json:"score"`
	Date       time.Time `json:"date"`
}
