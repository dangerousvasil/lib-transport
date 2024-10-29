package tags

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lib-transport/bgserver"
	"lib-transport/ptransport"
)

func NewTagsServer() *Server {
	return &Server{}
}

type Server struct {
	ptransport.UnsafeTagsServer
	bgserver.BaseService
}

func (s Server) List(ctx context.Context, request *ptransport.TagListRequest) (*ptransport.TagListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Get(ctx context.Context, request *ptransport.TagIdRequest) (*ptransport.TagResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Search(ctx context.Context, request *ptransport.TagSearchRequest) (*ptransport.TagListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Put(ctx context.Context, lite *ptransport.TagLite) (*ptransport.TagResponse, error) {
	userClaims, err := s.GetUserClaimsFromCTX(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, `Please login on a main page`)
	} else if userClaims == nil {
		return nil, status.Error(codes.Unauthenticated, `Please login on a main page`)
	}
	//TODO implement me
	panic("implement me")
}
