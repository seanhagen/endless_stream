package grpc

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// TLSConfig ...
type TLSConfig interface {
	apply(*internalConfig) error
}

// WithTLSConfig ...
func WithTLSConfig(certPath, keyPath string) TLSConfig {
	return tlsConfig{certPath, keyPath}
}

type tlsConfig struct {
	certPath string
	keyPath  string
}

// apply ...
func (tc tlsConfig) apply(conf *internalConfig) error {
	cert, err := tls.LoadX509KeyPair(tc.certPath, tc.keyPath)
	if err != nil {
		return fmt.Errorf("unable to build TLS key pair: %w", err)
	}

	conf.serverOptions[ServerOptionKeyTLS] = grpc.Creds(credentials.NewServerTLSFromCert(&cert))
	return nil
}

// WithMutualTLSConfig ...
func WithMutualTLSConfig(certPath, keyPath, rootCAPath string) TLSConfig { //nolint:revive,ireturn
	return nil
}
