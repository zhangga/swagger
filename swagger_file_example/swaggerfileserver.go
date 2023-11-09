package swaggerfileserver

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	cors            bool = true // 是否允许跨域
	readTimeout          = 10 * time.Second
	writeTimeout         = 10 * time.Second
	serviceRealHost      = "127.0.0.1:9000" // 服务真实的host，用于替换swagger.yaml中的host
)

func RunTestSwaggerFileServer() {
	serveMux := http.ServeMux{} // swagger file
	serveMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		// 处理openapi.yaml，加入servers配置
		if strings.HasSuffix(r.URL.Path, "openapi.yaml") {
			// 设置响应头
			w.Header().Set("Content-Type", "application/x-yaml")
			if cors {
				// 设置CORS头部信息
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			w.Write([]byte(getOpenApiYaml()))
			return
		}

		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		//p = path.Join("ratelimiter", p)

		if cors {
			// 设置CORS头部信息
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		http.ServeFile(w, r, p)
	})

	// 启动swagger服务
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		panic(fmt.Errorf("swagger file server listen failed: %w", err))
	}

	fmt.Printf("swagger file server listen on: http://%s/swagger/\n", listener.Addr().String())

	server := http.Server{
		Handler:      &serveMux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	go func() {
		if err = server.Serve(listener); err != nil {
			panic(fmt.Errorf("swagger file server serve failed: %w", err))
		}
	}()
}

func getOpenApiYaml() string {
	if len(serviceRealHost) == 0 {
		return OpenAPI
	}

	sb := strings.Builder{}
	sb.WriteString("servers:\n")
	sb.WriteString(fmt.Sprintf("  - url: http://%s\n", serviceRealHost))
	sb.WriteString("    description: 替换成服务接口的IP和说明\n")
	sb.WriteString(OpenAPI)
	return sb.String()
}
