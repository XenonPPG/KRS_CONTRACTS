package models

type Note struct {
	ID      int64  `gorm:"primary_key" json:"id"`
	Title   string `gorm:"not null" json:"title" validate:"required,min=1,max=100"`
	Content string `json:"content" validate:"max=5000"`

	UserID int64 `gorm:"not null" json:"user_id" validate:"required"`
}
