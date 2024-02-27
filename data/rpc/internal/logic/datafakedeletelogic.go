package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataFakeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDataFakeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataFakeDeleteLogic {
	return &DataFakeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DataFakeDeleteLogic) DataFakeDelete(in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteResponse{Successed: true}, nil
}
