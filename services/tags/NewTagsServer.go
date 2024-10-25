package tags

import (
	"context"
	"lib-transport/ptransport"
)

func NewTagsServer() *Server {
	return &Server{}
}

type Server struct {
	ptransport.UnsafeTagsServer
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
	//TODO implement me
	panic("implement me")
}
