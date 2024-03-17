package service

import (
	"app/config"
	"app/dto/request"
	"app/grpc/proto"
	"app/model"

	"gorm.io/gorm"
)

type shopService struct {
	db *gorm.DB
}

type ShopService interface {
	CreateShop(reqShop request.ShopRequest, profile *proto.Profile) (*model.Shop, error)
	CheckDuplicateShop(shopName string, profileId uint) (bool, error)
}

func (s *shopService) CreateShop(reqShop request.ShopRequest, profile *proto.Profile) (*model.Shop, error) {
	newShop := &model.Shop{
		Name:      reqShop.Name,
		Address:   reqShop.Address,
		ProfileID: uint(profile.ID),
	}

	if err := s.db.Model(&model.Shop{}).Create(&newShop).Error; err != nil {
		return nil, err
	}

	return newShop, nil
}

func (s *shopService) CheckDuplicateShop(shopName string, profileId uint) (bool, error) {
	var listShop []model.Shop

	if err := s.db.
		Model(&model.Shop{}).
		Where("profile_id = ? AND name = ?", profileId, shopName).
		Find(&listShop).Error; err != nil {
		return false, err
	}

	if len(listShop) > 0 {
		return true, nil
	}

	return false, nil
}

func NewShopService() ShopService {
	return &shopService{
		db: config.GetDB(),
	}
}
