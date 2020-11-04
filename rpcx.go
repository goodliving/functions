package functions

import (
	"fmt"
	"github.com/shima-park/agollo"
)

type RpcxAddr struct {
	ServiceAddr string
	ConsulAddr string
}

const (
	FlagRpcxPort = "rpcx.port"
	FlagRpcxConsulAddr = "rpcx.consul.addr"
)

func GetRpcxAddr() RpcxAddr {
	serviceAddr := fmt.Sprintf("%s:%s", GetHostIP(), agollo.Get(FlagRpcxPort))
	consulAddr := agollo.Get(FlagRpcxConsulAddr)
	return RpcxAddr{
		ServiceAddr: serviceAddr,
		ConsulAddr:  consulAddr,
	}
}