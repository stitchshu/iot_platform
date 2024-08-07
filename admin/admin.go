package main

import (
	"context"
	"flag"
	"fmt"
	"iot/user/rpc/types/user"
	"net/http"

	"iot/admin/internal/config"
	"iot/admin/internal/handler"
	"iot/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("token") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("StatusUnauthorized"))
				return
			}
			auth, err := ctx.RpcUser.Auth(context.Background(), &user.UserAuthRequest{Token: r.Header.Get("token")})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("StatusUnauthorized"))
				return
			}
			ctx.AuthUser = auth
			next(w, r)
		}
	})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
