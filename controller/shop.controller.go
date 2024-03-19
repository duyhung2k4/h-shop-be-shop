package controller

import (
	"app/config"
	"app/dto/request"
	"app/grpc/proto"
	"app/service"
	"app/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type shopController struct {
	utilsJWT      utils.JwtUtils
	clientProfile proto.ProfileServiceClient
	shopService   service.ShopService
}

type ShopController interface {
	GetShop(w http.ResponseWriter, r *http.Request)
	GetTypeProduct(w http.ResponseWriter, r *http.Request)
	CheckDuplicateShop(w http.ResponseWriter, r *http.Request)
	CreateShop(w http.ResponseWriter, r *http.Request)
	CreateTypeProduct(w http.ResponseWriter, r *http.Request)
}

func (c *shopController) CreateShop(w http.ResponseWriter, r *http.Request) {
	var shopReq request.ShopRequest

	if err := json.NewDecoder(r.Body).Decode(&shopReq); err != nil {
		badRequest(w, r, err)
		return
	}

	mapData, errMapData := c.utilsJWT.GetMapData(r)
	if errMapData != nil {
		handleError(w, r, errMapData, 401)
		return
	}

	profileID := uint(mapData["profile_id"].(float64))
	profile, errProfile := c.clientProfile.GetProfile(context.Background(), &proto.GetProfileReq{
		ProfileID: uint64(profileID),
	})

	if errProfile != nil {
		internalServerError(w, r, errProfile)
		return
	}

	newShop, errNewShop := c.shopService.CreateShop(shopReq, profile)
	if errNewShop != nil {
		internalServerError(w, r, errNewShop)
		return
	}

	res := Response{
		Data:    newShop,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *shopController) CheckDuplicateShop(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	shopName := queryValues.Get("name")

	mapData, errMapData := c.utilsJWT.GetMapData(r)
	if errMapData != nil {
		handleError(w, r, errMapData, 401)
		return
	}

	profileID := uint(mapData["profile_id"].(float64))
	isDuplicate, errCheckDuplicate := c.shopService.CheckDuplicateShop(shopName, profileID)
	if errCheckDuplicate != nil {
		internalServerError(w, r, errCheckDuplicate)
	}

	res := Response{
		Data:    isDuplicate,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *shopController) CreateTypeProduct(w http.ResponseWriter, r *http.Request) {
	var listTypeProduct []request.TypeProductRequest
	if err := json.NewDecoder(r.Body).Decode(&listTypeProduct); err != nil {
		badRequest(w, r, err)
		return
	}

	newTypeProducts, errTypeProduct := c.shopService.CreateTypeProduct(listTypeProduct)
	if errTypeProduct != nil {
		internalServerError(w, r, errTypeProduct)
		return
	}

	res := Response{
		Data:    newTypeProducts,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *shopController) GetShop(w http.ResponseWriter, r *http.Request) {
	mapData, errMapData := c.utilsJWT.GetMapData(r)
	if errMapData != nil {
		handleError(w, r, errMapData, 401)
		return
	}

	profileID := uint(mapData["profile_id"].(float64))
	shops, err := c.shopService.GetShop(profileID)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    shops,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *shopController) GetTypeProduct(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	shopId, errConvert := strconv.Atoi(queryValues.Get("shopId"))
	if errConvert != nil {
		badRequest(w, r, errConvert)
		return
	}

	typeProducts, err := c.shopService.GetTypeProduct(uint(shopId))
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    typeProducts,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewShopController() ShopController {
	clientProfile := proto.NewProfileServiceClient(config.GetConnProfileGRPC())
	return &shopController{
		utilsJWT:      utils.NewJwtUtils(),
		clientProfile: clientProfile,
		shopService:   service.NewShopService(),
	}
}
