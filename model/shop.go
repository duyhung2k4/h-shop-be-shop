package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ProfileID uint   `json:"profileId"`
	Name      string `json:"name"`
	Address   string `json:"address"`

	TypeProducts []TypeProduct `json:"typeProducts" gorm:"foreignKey:ShopID"`
}
