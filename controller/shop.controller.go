package controller

import (
	"app/config"
	"app/dto/request"
	"app/grpc/proto"
	"app/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

type shopController struct {
	utilsJWT      utils.JwtUtils
	clientProfile proto.ProfileServiceClient
}

type ShopController interface {
	CreateShop(w http.ResponseWriter, r *http.Request)
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

	res := Response{
		Data:    profile,
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
	}
}
