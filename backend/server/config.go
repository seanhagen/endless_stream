package server

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Config ...
type Config struct {
	Version            string
	Build              string
	Repo               string
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
}

func getConfig(cnf Config) (*viper.Viper, error) {
	vp := viper.New()
	vp.SetDefault("env", "development")

	vp.Set("version", cnf.Version)
	vp.Set("build", cnf.Build)
	vp.Set("repo", cnf.Repo)

	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("/etc/es")
	vp.AddConfigPath("$HOME/.es")
	vp.AddConfigPath(".")

	if err := vp.ReadInConfig(); err != nil {
		// Config file not found, or was found but another error was produced
		return nil, err
	}

	return vp, nil
}
