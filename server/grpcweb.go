package server

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// WrapServerForGrpcWeb wraps a grpc.Server in an http.Server that handles grpc-web requests (and also passes plain
// grpc to the wrapped server).
// When not the empty string, files in staticDirectory will be served for non-grpc requests.
// Remote grpc-web requests will pass CORS when the remote hostname (excluding port) equals allowedOriginHostname.
func WrapServerForGrpcWeb(grpcServer *grpc.Server, staticDirectory string, allowedOriginHostname string) *http.Server {
	grpcWebServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		originUrl, err := url.Parse(origin)
		if err != nil {
			return false
		}
		return originUrl.Hostname() == allowedOriginHostname
	}))
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
			if !strings.HasPrefix(r.URL.Path, "/api.Doko/") {
				static.ServeHTTP(w, r)
			}
			grpcWebServer.ServeHTTP(w, r)
		}), &http2.Server{}),
	}
	return httpServer
}
