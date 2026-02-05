package models

type User struct {
	ID        int64  `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	Name      string `gorm:"not null" json:"name" validate:"required,min=1,max=32"`
	Password  string `gorm:"not null" json:"password,omitempty" validate:"required,min=8"`
	DarkTheme bool   `gorm:"default:false" json:"dark_theme"`

	Notes []Note `gorm:"foreignkey:UserID" json:"notes,omitempty"`
}
