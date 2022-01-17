package rpcLib

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/lyj0309/jwc-lib/lib"
	etcdClient "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"log"
	"net/http"
	"time"
)

func NewRPC() *client.XClient {
	discovery, err := etcdClient.NewEtcdV3Discovery(
		"/rpc_jwc",
		"JWC",
		[]string{lib.Config.EtcdAddr},
		true,
		nil)

	if err != nil {
		log.Panicln("rpc初始化错误", err)
	}

	option := client.DefaultOption
	option.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	option.TCPKeepAlivePeriod = time.Hour * 24

	XClient := client.NewXClient("JWC", client.Failover, client.SelectByUser, discovery, option)
	XClient.Auth(lib.Config.RpcAuth)
	XClient.SetSelector(&roundRobinSelector{})
	log.Println("rpc初始化成功")
	return &XClient
}

type roundRobinSelector struct {
	servers []string
	i       int
}

func (s *roundRobinSelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {

	ss := s.servers
	if len(ss) == 0 {
		return ""
	}
	i := s.i
	i = i % len(ss)
	s.i = i + 1
	//fmt.Println(`所有服务器`, s.servers, `本次使用`, ss[i])
	return ss[i]
}

func (s *roundRobinSelector) UpdateServer(servers map[string]string) {
	ss := make([]string, 0, len(servers))
	for k := range servers {
		ss = append(ss, k)
	}
	if len(s.servers) != len(ss) {
		fmt.Println(`更新rpcserver`, servers)
		dataType, _ := json.Marshal(servers)
		dataString := string(dataType)
		http.Get(`https://sc.ftqq.com/SCU91220T6fe4d15407a2584b89fc1103ed037c0a6012ca6dc8057.send?text=` + dataString)
	}

	s.servers = ss
}
