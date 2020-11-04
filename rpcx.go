package functions

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/shima-park/agollo"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
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

func AddConsulRegistryPlugin(s *server.Server, basePath, addr, consulAddr string) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ConsulServers:  []string{consulAddr},
		BasePath:       basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}

	s.Plugins.Add(r)
}