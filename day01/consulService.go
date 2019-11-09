package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net"
	"net/rpc"
)

//1.定义一个服务
type UserLogin struct{}

//2.把这个服务实现
func (user *UserLogin) Login(req string, resp *string)error {
	*resp = req + "word"
	fmt.Println("功能完成")
	return nil
}
func main() {
	//初始化配置
	conConfig := api.DefaultConfig()
	//根据配置生成consul实例

	conClient, err := api.NewClient(conConfig)
	if err != nil {
		fmt.Println("根据配置生成consul", err)
		return
	}
	//实例化注册对象
	registerObj := api.AgentServiceRegistration{
		ID:      "1",
		Name:    "login",
		Port:    1234,
		Address: "192.168.12.37",
		Check: &api.AgentServiceCheck{
			CheckID:  "11",
			Name:     "login",
			TCP:      "192.168.12.37:1234",
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	//注册服务
	conClient.Agent().ServiceRegister(&registerObj)

	//注册服务
	rpc.RegisterName("login",new(UserLogin))

	//开启监听,建立连接
	listener,err:=net.Listen("tcp","192.168.12.37:1234")
	if err != nil {
		fmt.Println("net.Listen err:",err)
		return
	}
	defer listener.Close()
	//启动监听
	fmt.Println("建立连接 开始监听")

	for  {
		//fmt.Println("buu")
		conn,err:=listener.Accept()
		if err != nil {
			fmt.Println("建立连接失败:err",err)
			return
		}
		defer conn.Close()

		//在conn上绑定rpc
		rpc.ServeConn(conn)
	}

}
