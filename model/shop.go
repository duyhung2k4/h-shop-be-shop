package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ProfileID uint   `json:"profileId" gorm:"uniqueIndex:idx_profile_name"`
	Name      string `json:"name" gorm:"uniqueIndex:idx_profile_name"`
	Address   string `json:"address"`

	TypeProducts []TypeProduct `json:"typeProducts" gorm:"foreignKey:ShopID"`
}
