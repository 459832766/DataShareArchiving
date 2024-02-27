package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDataByPathLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListDataByPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDataByPathLogic {
	return &ListDataByPathLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListDataByPathLogic) ListDataByPath(in *pb.QueryRequest) (*pb.QueryResponse, error) {
	// todo: add your logic here and delete this line
	items := []*pb.DataItem{
		{Name: "request", Path: in.Path},
		{Name: "file1", Path: "path1"},
		{Name: "file2", Path: "path2"},
	}
	return &pb.QueryResponse{Items: items}, nil
}
