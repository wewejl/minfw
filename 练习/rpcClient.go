package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//conn,err:=rpc.Dial("tcp",":1234")
	conn,err:=net.Dial("tcp",":1234")
	if err != nil {
		fmt.Println("rpc.Dial err:",err)
		return
	}
	defer conn.Close()
	repConn:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//传入参数
	req:=100
	//传出参数
	var resp int
	//调用远程服务
	err=repConn.Call("hello.CallFunc",req,&resp)
	//打印数据
	fmt.Println("获取数据的为",resp)
}