package main

import (
	"context"
	"flag"
	"fmt"
	"lib-transport/bgctx"
	"lib-transport/jwtman"
	"lib-transport/pb"
	service "lib-transport/server"
	"lib-transport/services/tags"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

const (
	tokenDuration = 24 * 5 * time.Hour
)

func main() {
	//signal listner
	ctx := bgctx.InitCTX()

	port := flag.Int("port", 8082, "the server port")
	portGW := flag.Int("gw", 8081, "the gw port")
	JwtKey := flag.String("JwtKey", "lksdhfkjasjkfagsjkafgkjsfbj", "the JwtKey port")
	serverType := flag.String("type", "both", "type of server (grpc/rest/both)")
	endPoint := flag.String("endpoint", "", "gRPC endpoint")
	flag.Parse()

	if *endPoint == "" {
		*endPoint = fmt.Sprintf("127.0.0.1:%d", *port)
	}

	jwtManager := jwtman.NewJWTManager(*JwtKey)

	tagServer := tags.NewTagsServer()

	s := service.New(jwtManager, func(grpcServer *grpc.Server) {
		pb.RegisterTagsServer(grpcServer, tagServer)
	}, func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
		if err = pb.RegisterTagsHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
			return err
		}

		return nil
	})

	s.AddAccessibleRoles(tags.AccessibleRoles())

	// Start service
	s.Start(ctx, *serverType, *endPoint, *port, *portGW)
	slog.Info("Stop")
	<-ctx.Done()

	s.Stop()
	slog.Info("done")
}

type Configs struct {
	ScanStorageBucket string `mapstructure:"scan_storage_bucket"`
}
