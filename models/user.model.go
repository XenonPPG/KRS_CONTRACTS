package models

import "github.com/XenonPPG/KRS_CONTRACTS/gen/user_v1"

type User struct {
	ID         int64              `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	Login      string             `gorm:"not null;unique" json:"login" validate:"required,min=1,max=32"`
	Password   string             `gorm:"not null" json:"password,omitempty" validate:"required,min=8"`
	ColorTheme user_v1.ColorTheme `gorm:"default:AUTO" json:"color_theme"`

	Notes []Note `gorm:"foreignkey:UserID" json:"notes,omitempty"`
}
