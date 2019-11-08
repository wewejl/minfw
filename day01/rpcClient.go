package main

import "fmt"

func main() {
	//conn,err:=rpc.Dial("tcp",":1234")
	/*conn,err:=net.Dial("tcp",":1234")
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
	err=repConn.Call(ServiceName+"CallFunc",req,&resp)
	//打印数据
	fmt.Println("获取数据的为",resp)*/
	cli,err:=initClient(":1234")
	if err != nil {
		fmt.Println("客户端初始化错误:err",err)
		return
	}
	req :=10 //传入参数
	var resp int //传出参数
	//调用方法
	err=cli.CallFunc(req,&resp)
	if err != nil {
		fmt.Println("调用 cli.CallFunc :err",err)
		return
	}
	fmt.Println("得到数据resp是",resp)

}