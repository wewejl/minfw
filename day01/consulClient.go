package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net/rpc"
	"strconv"
)

func main() {
	//服务发现
	//初始化配置
	conConfig :=api.DefaultConfig()

	//获取consul实例
	conClient,err:=api.NewClient(conConfig)
	if err != nil {
		fmt.Println("初始化consul失败",err)
		return
	}


	//获取健康的服务地址
	serviceEntry,_,_:=conClient.Health().Service("login","",false,&api.QueryOptions{})


	//连接服务并判定rpc
	conn,err:=rpc.Dial("tcp",serviceEntry[0].Service.Address+":"+strconv.Itoa(serviceEntry[0].Service.Port))
	if err != nil {
		fmt.Println("连接服务 rpc err:",err)
		return
	}
	defer conn.Close()
	req :="hello" //传入参数
	var resp string //传出参数

	err=conn.Call("login.Login",req,&resp)
	if err != nil {
		fmt.Println("调用方法的时候出错误了",err)
		return
	}
	fmt.Println(resp)
}
