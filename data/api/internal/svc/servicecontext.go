package svc

import (
	"api/internal/config"
	"rpc/datarpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	DataRpc datarpc.Datarpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DataRpc: datarpc.NewDatarpc(zrpc.MustNewClient(c.DataRpc)),
	}
}
