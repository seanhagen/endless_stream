package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// ErrUnableToAppendCert ...
var ErrUnableToAppendCert = errors.New("unable to append CA certs to CA pool")

// TLSConfig ...
type TLSConfig interface {
	apply(*internalConfig) error
}

func WithInsecureTLS() TLSConfig {
	return insecureTLSConfig{}
}

type insecureTLSConfig struct{}

// apply ...
func (itls insecureTLSConfig) apply(conf *internalConfig) error {
	conf.serverOptions[ServerOptionKeyTLS] = grpc.Creds(insecure.NewCredentials())
	return nil
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
		return fmt.Errorf("unable to build TLS key pair from file: %w", err)
	}

	tlsConf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
		MinVersion:   tls.VersionTLS12,
	}

	conf.serverOptions[ServerOptionKeyTLS] = grpc.Creds(credentials.NewTLS(tlsConf))
	return nil
}

// WithMutualTLSConfig ...
func WithMutualTLSConfig(certPath, keyPath, rootCAPath string) TLSConfig { //nolint:revive,ireturn
	return mutualTLSConfig{certPath: certPath, keyPath: keyPath, rootPath: rootCAPath}
}

type mutualTLSConfig struct {
	certPath, keyPath, rootPath string
}

// apply ...
func (mtls mutualTLSConfig) apply(conf *internalConfig) error {
	cert, err := tls.LoadX509KeyPair(mtls.certPath, mtls.keyPath)
	if err != nil {
		return fmt.Errorf("unable to build TLS key pair from file: %w", err)
	}

	data, err := os.ReadFile(mtls.rootPath)
	if err != nil {
		return fmt.Errorf("unable to read CA certs: %w", err)
	}

	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(data) {
		return ErrUnableToAppendCert
	}

	tlsConf := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    capool,
		MinVersion:   tls.VersionTLS12,
	}

	conf.serverOptions[ServerOptionKeyTLS] = grpc.Creds(credentials.NewTLS(tlsConf))
	return nil
}
