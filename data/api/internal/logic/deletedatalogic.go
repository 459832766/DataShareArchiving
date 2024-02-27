package logic

import (
	"context"
	"rpc/pb"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDataLogic {
	return &DeleteDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDataLogic) DeleteData(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	// todo: add your logic here and delete this line
	in := &pb.DeleteRequest{
		Key:  "file1",
		Path: "path1",
	}
	pbResponse, _ := l.svcCtx.DataRpc.DataFakeDelete(l.ctx, in)
	return &types.DeleteResponse{Successed: pbResponse.Successed}, nil
}
