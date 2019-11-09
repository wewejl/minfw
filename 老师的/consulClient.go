package main

import (
	"net/rpc"
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
)

func main(){
	//服务发现
	//初始化配置
	conConfig := api.DefaultConfig()

	//获取consul实例
	conClient ,err:= api.NewClient(conConfig)
	if err != nil {
		fmt.Println("初始化consul失败",err)
		return
	}

	//服务的注销
	conClient.Agent().ServiceDeregister("1")

	//获取健康的服务地址
	serviceEntry,_,_ :=conClient.Health().Service("login","",false,&api.QueryOptions{})

	//负载均衡

	//链接服务并绑定rpc
	conn,err := rpc.Dial("tcp",serviceEntry[0].Service.Address+":"+strconv.Itoa(serviceEntry[0].Service.Port))
	if err != nil {
		fmt.Println("建立链接失败",err)
		return
	}
	defer conn.Close()
	req := "hello"
	var resp string

	err = conn.Call("login.Login",req,&resp)

	fmt.Println(resp)
}
