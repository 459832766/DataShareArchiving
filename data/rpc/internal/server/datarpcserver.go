// Code generated by goctl. DO NOT EDIT.
// Source: datarpc.proto

package server

import (
	"context"

	"rpc/internal/logic"
	"rpc/internal/svc"
	"rpc/pb"
)

type DatarpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedDatarpcServer
}

func NewDatarpcServer(svcCtx *svc.ServiceContext) *DatarpcServer {
	return &DatarpcServer{
		svcCtx: svcCtx,
	}
}

func (s *DatarpcServer) DataFakeDelete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	l := logic.NewDataFakeDeleteLogic(ctx, s.svcCtx)
	return l.DataFakeDelete(in)
}

func (s *DatarpcServer) ListDataByPath(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	l := logic.NewListDataByPathLogic(ctx, s.svcCtx)
	return l.ListDataByPath(in)
}
