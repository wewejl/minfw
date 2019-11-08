package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//1.首先要有一个结构体
type HelloWorld struct {
	
}

//2.定义一个方法, 第一个参数为传入参数,第二个参数为传出参数    引用传递组 返回值 有 只能有一个error
func (h *HelloWorld)CallFunc(req int,resp *int) error {
	*resp=req+22
	return nil
}

//3.主函数
func main()  {
	//注册服务     本质  在内存维护了一张哈希表
	//rpc.RegisterName(ServiceName,new(HelloWorld))
	RegisterService(new(HelloWorld))//这里地方就实现

	//创建监听连接
	listener,err:=net.Listen("tcp",":1234")
	if err != nil {
		fmt.Println("设置监听失败",err)
		return
	}
	fmt.Println("开启监听..")
	//建立连接
	conn,err:=listener.Accept()
	if err != nil {
		fmt.Println("建立连接失败",err)
		return
	}
	defer conn.Close()

	//用rpc连接
	//rpc.ServeConn(conn)
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}