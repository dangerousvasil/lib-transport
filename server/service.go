package service

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"lib-transport/interceptor"
	"lib-transport/jwtman"
	"log"
	"net"
	"net/http"
	"sync"
)

type registerRESTServers func(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) (err error)

type registerGRPCServers func(server *grpc.Server)

func New(manager *jwtman.JWTManager, registerGRPCServersFc registerGRPCServers, registerRESTServersFc registerRESTServers) *Service {
	return &Service{
		wg:                  new(sync.WaitGroup),
		jwtManager:          manager,
		registerGRPCServers: registerGRPCServersFc,
		registerRESTServers: registerRESTServersFc,
		accessibleRoles:     make(map[string][]string),
	}
}

func (s *Service) AddAccessibleRoles(accessibleRoles map[string][]string) {
	for path, roles := range accessibleRoles {
		s.accessibleRoles[path] = roles
	}
}

type Service struct {
	wg                  *sync.WaitGroup
	jwtManager          *jwtman.JWTManager
	registerGRPCServers registerGRPCServers // ф-ия для RegisterServer
	registerRESTServers registerRESTServers // ф-ия для RegisterServiceHandlerFromEndpoint
	accessibleRoles     map[string][]string // TODO: подумать как это сделать более изящное задание доступов
}

func (s *Service) Start(ctx context.Context, serverType, endPoint string, port int, portGW int) {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	addressGW := fmt.Sprintf("0.0.0.0:%d", portGW)

	if serverType == "grpc" || serverType == "both" {
		go func() {
			err := s.grpcServer(ctx, address)
			if err != nil {
				log.Fatal("cannot start server: ", err)
			}
		}()
		log.Printf("Started GRPC server at %s", address)
	}

	if serverType == "rest" || serverType == "both" {
		listenerGW, err := net.Listen("tcp", addressGW)
		if err != nil {
			log.Fatal("cannot listener GW server: ", err)
		}
		err = s.runRESTServer(ctx, listenerGW, endPoint)
		if err != nil {
			log.Fatal("cannot start GW server: ", err)
		}

		log.Printf("Started GW GRPC server at %s", addressGW)
	}
}

func (s *Service) Stop() {
	s.wg.Wait()
}

func (s *Service) grpcServer(ctx context.Context, address string) error {
	s.wg.Add(1)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	interceptor := interceptor.NewAuthInterceptor(s.jwtManager, s.accessibleRoles)
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	grpcServer := grpc.NewServer(serverOptions...)

	s.registerGRPCServers(grpcServer)

	reflection.Register(grpcServer)

	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatal("cannot start server: ", err)
		}
	}()
	<-ctx.Done()

	grpcServer.GracefulStop()

	s.wg.Done()
	return nil
}

func (s *Service) runRESTServer(ctx context.Context, listener net.Listener, grpcEndpoint string) error {

	mux := runtime.NewServeMux()

	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := s.registerRESTServers(ctx, mux, grpcEndpoint, dialOptions); err != nil {
		log.Println("register")
		return err
	}
	if err := mountSwagger(mux); err != nil {
		return err
	}
	if err := mountProtoc(mux); err != nil {
		return err
	}

	go func() {
		err := http.Serve(listener, mux)
		if err != nil {
			return
		}
	}()

	return nil
}
