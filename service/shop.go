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
	GetShop(profileId uint) ([]model.Shop, error)
	GetTypeProduct(shopId uint) ([]model.TypeProduct, error)
	CreateShop(reqShop request.ShopRequest, profile *proto.Profile) (*model.Shop, error)
	CheckDuplicateShop(shopName string, profileId uint) (bool, error)
	CreateTypeProduct(data []request.TypeProductRequest) ([]model.TypeProduct, error)
}

func (s *shopService) GetShop(profileId uint) ([]model.Shop, error) {
	var shops []model.Shop

	if err := s.db.
		Model(&model.Shop{}).
		Where("profile_id = ?", profileId).
		Find(&shops).Error; err != nil {
		return []model.Shop{}, err
	}

	return shops, nil
}

func (s *shopService) GetTypeProduct(shopId uint) ([]model.TypeProduct, error) {
	var typeProducts []model.TypeProduct

	if err := s.db.
		Model(&model.TypeProduct{}).
		Where("id = ?", shopId).
		Find(&typeProducts).Error; err != nil {
		return []model.TypeProduct{}, err
	}

	return typeProducts, nil
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

func (s *shopService) CreateTypeProduct(data []request.TypeProductRequest) ([]model.TypeProduct, error) {
	var listNewProduct []model.TypeProduct

	for _, p := range data {
		listNewProduct = append(listNewProduct, model.TypeProduct{
			ShopID: p.ShopID,
			Hastag: p.Hastag,
			Name:   p.Name,
		})
	}

	if err := s.db.Model(&model.TypeProduct{}).Create(&listNewProduct).Error; err != nil {
		return []model.TypeProduct{}, err
	}

	return listNewProduct, nil
}

func NewShopService() ShopService {
	return &shopService{
		db: config.GetDB(),
	}
}
