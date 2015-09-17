package ipc

import (
	"encoding/json"
	"fmt"
	"net"
)


type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct{
	Server
	Addr string
}

func NewIpcServer(server Server) *IpcServer{
	return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string){
		for{
			request := <-c

			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil{
				fmt.Println("Invalid request format:", request)
			}

			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b)

		}

		fmt.Print("Session closed.")
	}(session)

	fmt.Print("A new session has been created.")
	return session
}

//

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (srv *IpcServer) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{l.(*net.TCPListener)})
}

func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	return nil
}