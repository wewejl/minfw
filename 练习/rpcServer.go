package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//1.创建一个结构体
type Helloword struct {
	
}

//2.创建一个方法 里面要有传入参数 和传出参数
func (h *Helloword)callword(rep int,resp *int) error {
	*resp+=rep
	return nil
}

func main() {
	//注册服务     本质  在内存维护了一张哈希表
	err:=rpc.RegisterName("hello",new(Helloword))
	if err != nil {
		fmt.Println("rpc.RegisterName err",err)
		return
	}
	linster,err:=net.Listen("tcp",":1234")
	if err != nil {
		fmt.Println("Listen err:",err)
		return
	}
	defer linster.Close()
	conn,err:=linster.Accept()
	if err != nil {
		fmt.Println("",err)
		return
	}

	//rpc.ServeConn(coon)
	//
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}
