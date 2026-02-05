package models

type Note struct {
	ID      int    `gorm:"primary_key" json:"id"`
	Name    string `gorm:"not null" json:"name" validate:"required,min=1,max=100"`
	Content string `json:"content" validate:"max=5000"`

	UserID int `gorm:"not null" json:"user_id" validate:"required"`
}
