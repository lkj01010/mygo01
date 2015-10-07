package main

//import(
//	"fmt"
//	"math"
//)
//
//func Squrt(i int) int{
//	v := math.Sqrt(float64(i))
//	return int(v)
//}
//
//func main() {
//	fmt.Println("result is ", Squrt(23))
//}

//import (
//    "fmt"
//    "reflect"
//)

//func main() {
//    type S struct {
//         F string `species:"gopher" color:"blue"`
//     }
//
//     s := S{}
//     st := reflect.TypeOf(s)
//     field := st.Field(0)
//     fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
//
// }

import (
	"fmt"
	"net"
)

const (
	ip   = ""
	port = 3333
)

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待客户端连接...")
	Server(listen)
}

func Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close()
		go func() {
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data)
				fmt.Println("客户端发来数据:", string(data[0:i]))
				if err != nil {
					fmt.Println("读取客户端数据错误:", err.Error())
					break
				}
				outstr := "finish_hahah"
//				conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
				conn.Write([]byte(outstr))
			}

		}()
	}
}
