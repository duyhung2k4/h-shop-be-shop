package model

import "gorm.io/gorm"

type TypeProduct struct {
	gorm.Model
	ShopID uint   `json:"shopId" gorm:"uniqueIndex:idx_shop_hastag"`
	Hastag string `json:"hastag" gorm:"uniqueIndex:idx_shop_hastag"`
	Name   string `json:"name"`

	Shop *Shop `json:"shop" gorm:"foreignKey:ShopID"`
}
