package main

import (
	"fmt"
	"net/rpc"
)

func main()  {
	// 连接服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()
	// 调用远程函数
	var reply string
	err = conn.Call("hello.HelloWorld", "大哥", &reply)
	if err != nil {
		fmt.Println("Call :", err)
		return
	}
	fmt.Println(reply)
}
