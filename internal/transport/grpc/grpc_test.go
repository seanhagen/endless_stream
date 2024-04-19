package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGrpc_Constructor(t *testing.T) {
	conf := Config{}

	transport, err := New(conf)

	assert.NoError(t, err)
	require.NotNil(t, transport)
	assert.IsType(t, (*Transport)(nil), transport)
}
