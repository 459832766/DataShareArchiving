package logic

import (
	"context"
	"net/http"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadDataLogic {
	return &UploadDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadDataLogic) UploadData(r *http.Request) (resp *types.UploadResponse, err error) {
	file, _, err := r.FormFile("file")

	if err != nil || file == nil {
		return
	}

	defer file.Close()

	return &types.UploadResponse{Path: "fake path."}, nil
}
