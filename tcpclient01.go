package main

import (
	"fmt"
	"net"
)

const (
	addr = "127.0.0.1:3333"
)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
//	Client(conn)
	Client2(conn)
}

func Client(conn net.Conn) {
	sms := make([]byte, 128)
	for {
		fmt.Print("请输入要发送的消息:")
		_, err := fmt.Scan(&sms)
		if err != nil {
			fmt.Println("数据输入异常:", err.Error())
		}
		conn.Write(sms)
		buf := make([]byte, 128)
		c, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
		}
		fmt.Println(string(buf[0:c]))
	}

}

func Client2(conn net.Conn){
	sms := make([]byte, 128)
	c := make(chan string)
	fmt.Println("1")
	go func(sms_ *[]byte, c_ chan string){
		fmt.Println("11")
		for{
			_, err:= fmt.Scan(sms_)
			if err != nil{
				fmt.Println("数据输入异常:", err.Error())
			}
			fmt.Println("scan enter: ", string(*sms_))
			c_ <- string(*sms_)

		}
	}(&sms, c)

	fmt.Println("2")
	go func(c chan string, conn net.Conn){
		fmt.Println("21")
		buf := make([]byte, 128)
		for{
//			smsSlice := make([]byte, 100)
			fmt.Println("s := <- c0")
			s := <- c
			fmt.Println("s := <- c0")
//			outStr := []{"client send"} + sms
//			fmt.Println(outStr)
//			strSlice := outStr[0:10]
			smsSlice := []byte(s)
			conn.Write(smsSlice)
			_, err := conn.Read(buf)
			if err != nil{
				fmt.Println("读取服务器数据异常:", err.Error())
			}
			fmt.Println(string(buf))
		}
	}(c, conn)

	flag := make(chan interface{})
	<- flag
}