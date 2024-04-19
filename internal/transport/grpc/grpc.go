// Package grpc ...
package grpc

import (
	"context"

	"github.com/seanhagen/endless_stream/internal/proto"
)

// HexServer ...
type HexServer interface {
	Info(context.Context, *proto.InfoRequest) (*proto.InfoResponse, error)
	Game(proto.Hex_GameServer) error
}

// AdminServer ...
type AdminServer interface {
	Manage(proto.Admin_ManageServer) error
}

// ServerList ...
type ServerList struct {
	Hex   HexServer
	Admin AdminServer
}

// Config ...
type Config struct {
	Servers ServerList
}

// Transport ...
type Transport struct {
	proto.UnimplementedHexServer
	proto.UnimplementedAdminServer
}

// New ...
func New(_ Config) (*Transport, error) {
	return &Transport{}, nil
}
