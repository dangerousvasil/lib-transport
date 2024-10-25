package tags

import (
	"context"
	"lib-transport/transport"
)

func NewTagsServer() *Server {
	return &Server{}
}

type Server struct {
	transport.UnsafeTagsServer
}

func (s Server) List(ctx context.Context, request *transport.TagListRequest) (*transport.TagListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Get(ctx context.Context, request *transport.TagIdRequest) (*transport.TagResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Search(ctx context.Context, request *transport.TagSearchRequest) (*transport.TagListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Put(ctx context.Context, lite *transport.TagLite) (*transport.TagResponse, error) {
	//TODO implement me
	panic("implement me")
}
