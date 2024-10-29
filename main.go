package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"lib-transport/bgctx"
	"lib-transport/bgserver"
	"lib-transport/jwtman"
	"lib-transport/ptransport"
	"lib-transport/services/tags"
	"log"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {
	//signal listner
	ctx := bgctx.InitCTX()
	JwtKey := flag.String("JwtKey", "lksdhfkjasjkfagsjkafgkjsfbj", "the jwt key")

	port := flag.Int("port", 8080, "the server port (0 - disable)")
	portGW := flag.Int("gw", 8082, "the gateway api port (0 - disable)")
	endPoint := flag.String("endpoint", "", "gRPC endpoint for remote server")

	serverType := ""
	if *port > 0 && *portGW > 0 {
		serverType = "both"
		if *endPoint != "" {
			log.Fatal("remote grpc endpoint must be empty")
		}
	} else if *port > 0 && *portGW == 0 {
		serverType = "grpc"
	} else if *port == 0 && *portGW > 0 {
		serverType = "rest"
		if *endPoint == "" {
			log.Fatal("remote grpc endpoint must be set")
		}
	}

	if *endPoint == "" {
		*endPoint = fmt.Sprintf("127.0.0.1:%d", *port)
	}

	jwtManager := jwtman.NewJWTManager(*JwtKey)

	tagServer := tags.NewTagsServer()

	s := bgserver.New(jwtManager, func(grpcServer *grpc.Server) {
		ptransport.RegisterTagsServer(grpcServer, tagServer)
	}, func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
		if err = ptransport.RegisterTagsHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
			return err
		}
		return nil
	})

	s.AddAccessibleRoles(tags.AccessibleRoles())
	s.AddStatic(true)

	// Start service
	s.Start(ctx, serverType, *endPoint, *port, *portGW)
	slog.Info("Stop")
	<-ctx.Done()

	s.Stop()
	slog.Info("done")
}
