package main

import (
	"flag"
	"fmt"
	"net/http"

	"api_jwt/internal/config"
	"api_jwt/internal/handler"
	"api_jwt/internal/svc"

	"github.com/fengfengzhidao/go-zero/common/response"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// JwtUnauthorizedResult jwt验证失败的回调
func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err) // 具体的错误，没带token，token过期？伪造token？
	httpx.WriteJson(w, http.StatusOK, response.Body{10087, nil, "鉴权失败"})
}
