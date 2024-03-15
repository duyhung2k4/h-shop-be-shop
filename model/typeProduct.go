package model

import "gorm.io/gorm"

type TypeProduct struct {
	gorm.Model
	ShopID uint   `json:"shopId"`
	Hastag string `json:"hastag"`
	Name   string `json:"name"`

	Shop *Shop `json:"shop" gorm:"foreignKey:ShopID"`
}
