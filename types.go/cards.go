package types

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	CardType string `gorm:"cardType" json:"card_type,omitempty"`
	Number   string `gorm:"number" json:"number,omitempty"`
	UserID   uint   `gorm:"userid" json:"user_id,omitempty"`
}
