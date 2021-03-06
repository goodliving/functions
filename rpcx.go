package functions

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/shima-park/agollo"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"net"
	"time"
)

type RpcxInfo struct {
	AppName string
	ServiceAddr string
	ConsulAddr string
	RpcxBasePath string
}

const (
	FlagAppName = "app.name"
	FlagRpcxPort = "rpcx.port"
	FlagRpcxConsulAddr = "rpcx.consul.addr"
	FlagRpcxBasePath = "rpcx.base.path"
)

func GetRpcxInfo() RpcxInfo {
	appName := agollo.Get(FlagAppName)
	serviceAddr := fmt.Sprintf("%s:%s", GetHostIP(), agollo.Get(FlagRpcxPort))
	consulAddr := agollo.Get(FlagRpcxConsulAddr)
	basePath := agollo.Get(FlagRpcxBasePath)
	return RpcxInfo{
		AppName: appName,
		ServiceAddr: serviceAddr,
		ConsulAddr:  consulAddr,
		RpcxBasePath: basePath,
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
	s.Plugins.Add(&ConnectionPlugin{})
}

type ConnectionPlugin struct {
}


func (p *ConnectionPlugin) HandleConnAccept(conn net.Conn) (net.Conn, bool) {
	log.Printf("client %v connected", conn.RemoteAddr().String())
	return conn, true
}

func (p *ConnectionPlugin) HandleConnClose(conn net.Conn) bool {
	log.Printf("client %v closed", conn.RemoteAddr().String())
	return true
}