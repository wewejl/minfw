package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const ServiceName ="heeloo"

//当父类
type BaseService interface {
	CallFunc(int,*int)error
}

func RegisterService(service BaseService)  {
	rpc.RegisterName(ServiceName,service)
}

//客户端这个方法优化
type rpcClient struct {
	c *rpc.Client
}

//初始化 客户端
func initClient(addr string) (rpcClient,error) {
	conn,err:=net.Dial("tcp",addr)
	if err != nil {
		fmt.Println("net.Dial err:",err)
		return rpcClient{},err
	}
	repConn:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return rpcClient{c:repConn},nil
}

//调用方法
func (cli rpcClient)CallFunc(req int,resp *int) error {
	return cli.c.Call(ServiceName+".CallFunc",req,resp)
}


