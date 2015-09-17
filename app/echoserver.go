package main

import(
	"mygo01/ipc"
	"net"
	"fmt"
	"net/http"
)
type EchoServer struct{

}

func (server *EchoServer)Handle(method, params string) *ipc.Response{
	return &ipc.Request{"ok", "echo: " + method + "~" + params}
}

func (server *EchoServer)Name() string {
	return "EchoServer"
}

func Test01(){
	server := ipc.NewIpcServer(&EchoServer{})

	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	checkError(err)



}

func checkError(err error){
	if err != nil{
		fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}