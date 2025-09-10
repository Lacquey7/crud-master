package model

type Orders struct {
	Id            int    `gorm:"primaryKey,AutoIncrement" json:"id,omitempty"`
	UserId        string `json:"user_id"`
	NumberOfItems string `json:"number_of_items"`
	TotalAmount   string `json:"total_amount"`
}
