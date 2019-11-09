package main

import (
	"net/rpc"
	"net"
	"fmt"
	"github.com/hashicorp/consul/api"
)

//先定义一个服务
type UserLogin struct {}

func (this*UserLogin)Login(req string,resp*string)error{
	*resp = req + " world"
	return nil
}

func main(){
	//注册服务到consul上
	consulConfig := api.DefaultConfig()

	//根据配置获取consul实例操作
	conClient,err := api.NewClient(consulConfig)
	if err != nil{
		fmt.Println("获取consul实例失败",err)
		return
	}

	//实例化注册对象
	registerObj := api.AgentServiceRegistration{
		ID:"1",
		Name:"login",
		Address:"192.168.137.82",
		Port:1234,
		Check:&api.AgentServiceCheck{
			CheckID:"11",
			Name:"login",
			TCP:"192.168.137.82:1234",
			Interval:"5s",
			Timeout:"1s",
		},
	}

	//注册服务
	conClient.Agent().ServiceRegister(&registerObj)



	//注册服务
	rpc.RegisterName("login",new(UserLogin))

	//开启监听,建立链接
	listener ,err := net.Listen("tcp","192.168.137.82:1234")
	if err != nil {
		fmt.Println("建立监听失败")
		return
	}
	defer listener.Close()

	fmt.Println("开启监听...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立链接失败")
			return
		}
		defer conn.Close()

		//在conn上绑定rpc
		rpc.ServeConn(conn)
	}
}
