package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"name" json:"name,omitempty"`
	Gender string `gorm:"gender" json:"gender,omitempty"`
	Age    string `gorm:"age" json:"age,omitempty"`
}
