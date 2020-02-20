package grpc

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/websocket"
)

func shouldEnv(name, def string) string {
	if e := strings.TrimSpace(os.Getenv(name)); e != "" {
		return e
	}
	return def
}

func hstsHandler(fn http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("hsts handler: %v", r.URL.String())
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		fn(w, r)
	})
}

func grpcTrafficSplitter(fallback http.HandlerFunc, grpcHandler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("grpc traffic splitter received request: %v", r.URL.String())
		// Redirect gRPC and gRPC-Web requests to the gRPC Server

		if r.ProtoMajor == 2 || websocket.IsWebSocketUpgrade(r) {
			log.Printf("should be grpc")
			grpcHandler.ServeHTTP(w, r)
		} else {
			log.Printf("fallback")
			fallback(w, r)
		}
	})
}

// func fallback(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("request received by fallback handler")
// 	w.Write([]byte("ok"))
// }

type fallback struct{}

// ServeHTTP ...
func (fb fallback) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("request received by fallback handler")
	spew.Dump(r)
	w.Write([]byte("ok"))
}
