package grpc

import (
	"os"
	"strings"
)

func shouldEnv(name, def string) string {
	if e := strings.TrimSpace(os.Getenv(name)); e != "" {
		return e
	}
	return def
}
