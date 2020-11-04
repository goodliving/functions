package functions

import (
	"fmt"
	"github.com/shima-park/agollo"
)

type RpcxAddr struct {
	ServiceAddr string
	ConsulAddr string
}

func GetRpcxAddr() RpcxAddr {
	serviceAddr := fmt.Sprintf("%s:%s", GetHostIP(), agollo.Get("rpcx.port"))
	consulAddr := agollo.Get("rpcx.consul.addr")
	return RpcxAddr{
		ServiceAddr: serviceAddr,
		ConsulAddr:  consulAddr,
	}

}