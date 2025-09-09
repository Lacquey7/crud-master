package model

type Movie struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
