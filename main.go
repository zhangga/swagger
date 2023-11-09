package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/zhangga/swagger/swagger"
	"github.com/zhangga/swagger/version"
)

var (
	port     uint
	yamlHost string
)

func init() {
	flag.UintVar(&port, "port", 8081, "port to listen on")
	flag.StringVar(&yamlHost, "yaml", "localhost:9000", "host of the swagger yaml file")
}

func main() {
	flag.Parse()
	fmt.Println(version.Print())

	swagger.SwaggerHost = yamlHost

	serveMux := http.ServeMux{}

	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	serveMux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("version: " + version.Print()))
	})

	// swagger ui
	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	//serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.Handle(prefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 修正prefix
		if index := strings.LastIndex(r.URL.Path, "/"); index != -1 {
			prefix = r.URL.Path[:index]
		}
		// 使用swagger-ui的url指定yaml的host必须以'/'结尾
		if start := strings.Index(r.URL.Path, "/swagger-ui/"); start != -1 && strings.HasSuffix(r.URL.Path, "/") {
			begin := start + len("/swagger-ui/")
			end := strings.LastIndex(r.URL.Path, "/")
			if end > begin {
				yamlHost := r.URL.Path[begin:end]
				swagger.SwaggerHost = yamlHost
			}
		}
		http.StripPrefix(prefix, fileServer).ServeHTTP(w, r)
	}))

	// 启动swagger服务
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Errorf("swagger server listen failed: %w", err))
	}

	fmt.Printf("swagger server listen on: http://%s/swagger-ui\n", listener.Addr().String())

	server := http.Server{
		Handler:      &serveMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err = server.Serve(listener); err != nil {
		panic(fmt.Errorf("swagger server serve failed: %w", err))
	}
}
