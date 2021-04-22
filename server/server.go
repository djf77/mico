package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义类对象
type World struct {
}

// 绑定类方法
func (the *World) HelloWorld (name string, resp *string) error  {
	*resp = name + " 你好！"
	return nil
}

func main()  {
	// 注册rpc服务
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册rpc服务失败", err)
		return
	}
	// 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("开始监听")
	// 建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")
	// 绑定服务
	rpc.ServeConn(conn)
}