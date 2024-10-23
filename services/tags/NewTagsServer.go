package tags

import (
	"context"
	"lib-transport/pb"
)

func NewTagsServer() *Server {
	return &Server{}
}

type Server struct {
	pb.UnsafeTagsServer
}

func (s Server) List(ctx context.Context, request *pb.TagListRequest) (*pb.TagListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Get(ctx context.Context, request *pb.TagIdRequest) (*pb.TagResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Search(ctx context.Context, request *pb.TagSearchRequest) (*pb.TagListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Put(ctx context.Context, lite *pb.TagLite) (*pb.TagResponse, error) {
	//TODO implement me
	panic("implement me")
}
