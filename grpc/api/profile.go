package api

import (
	"app/grpc/proto"
	"context"
)

type profileServiceGRPC struct {
	proto.UnsafeProfileServiceServer
}

func (g *profileServiceGRPC) GetProfile(context.Context, *proto.GetProfileReq) (*proto.Profile, error) {
	return nil, nil
}

func NewProfileServiceGPRC() proto.ProfileServiceServer {
	return &profileServiceGRPC{}
}
