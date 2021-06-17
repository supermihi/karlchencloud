package server

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strings"
)

// https://rogchap.com/2019/07/26/in-process-grpc-web-proxy/
func WrapServer(grpcServer *grpc.Server, staticDirectory string) *http.Server {
	grpcWebServer := grpcweb.WrapServer(grpcServer)
	var static http.Handler
	if staticDirectory != "" {
		log.Printf("serving static files at %s", staticDirectory)
		staticRoot := http.Dir(staticDirectory)
		static = http.FileServer(staticRoot)
	} else {
		static = http.NotFoundHandler()
	}
	httpServer := &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
			if !strings.HasPrefix(r.URL.Path, "/api.Doko/") {
				static.ServeHTTP(w, r)
			}
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	return httpServer
}
