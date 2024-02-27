package logic

import (
	"context"
	"rpc/pb"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDataLogic {
	return &QueryDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDataLogic) QueryData(req *types.QueryRequest) (resp *types.QueryResponse, err error) {
	// todo: add your logic here and delete this line
	pbReq := pb.QueryRequest{Path: req.Path}
	pbResp, _ := l.svcCtx.DataRpc.ListDataByPath(l.ctx, &pbReq)
	pbItems := make([]types.DataItem, 0, len(pbResp.Items))
	for _, v := range pbResp.Items {
		pbItems = append(pbItems, types.DataItem{
			Path: v.Path,
			Name: v.Name,
		})
	}
	return &types.QueryResponse{Items: pbItems}, nil
}
