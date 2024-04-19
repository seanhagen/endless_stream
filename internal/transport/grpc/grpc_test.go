package grpc

import (
	"net"
	"testing"

	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
)

const grpcListen = ":10000"

func TestGrpc_Constructor(t *testing.T) {
	conf := Config{}

	transport, err := New(conf)

	assert.NoError(t, err)
	require.NotNil(t, transport)
	assert.IsType(t, (*Transport)(nil), transport)
}

func TestGrpc_Transport(t *testing.T) {
	controller := gomock.NewController(t)
	gameServer := proto.NewMockHex_GameServer(controller)
	adminServer := proto.NewMockAdminServer(controller)

	conf := Config{
		Servers: ServerList{
			Hex:   gameServer,
			Admin: adminServer,
		},
	}

	l, err := net.Listen("tcp", grpcListen)
	require.NoError(t, err)
	s := grpc.NewServer()
	proto.RegisterHexServer(s, srv)

	dopts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	// setup grpc dialer
	conn, err := grpc.Dial(grpcListen, dopts...)
	require.NoError(t, err, "unable to dial")

	transport, err := New(conf)

	client := proto.NewHexClient(cc)
}
