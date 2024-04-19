package grpc

import (
	"context"

	"github.com/seanhagen/endless_stream/internal/proto"
)

type HexServer interface {
	Info(context.Context, *proto.InfoRequest) (*proto.InfoResponse, error)
	Game(proto.Hex_GameServer) error
}

type AdminServer interface {
	Manage(proto.Admin_ManageServer) error
}

type ServerList struct {
	Hex   HexServer
	Admin AdminServer
}

type Config struct {
	Servers ServerList
}

type Transport struct {
	proto.UnimplementedHexServer
	proto.UnimplementedHexServer
}

func New(conf Config) (*Transport, error) {
	return &Transport{}, nil
}
