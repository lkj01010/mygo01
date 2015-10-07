package net

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
	"sync"
	"os"
)


type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type iServer interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct{
	iServer
	Addr string
}

func NewIpcServer(server iServer) *IpcServer{
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

/////////////////////////////////////////
type conn struct{
	remoteAddr string
	server     *Server
	rwc        net.Conn

	mu           sync.Mutex
	closeNotifyc chan bool
}

func (c *conn)serve(){
	origConn := c.rwc // copy it before it's set nil on Close or Hijack
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
	}
}


///////////////////////////////////////
type Server struct{
	Addr string
}
//func (srv *Server)
//

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func (srv *Server) ListenAndServe() error {
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

// Create new connection from rwc.
func (srv *Server) newConn(rwc net.Conn) (c *conn, err error) {
	c = new(conn)
	c.remoteAddr = rwc.RemoteAddr().String()
	c.server = srv
	c.rwc = rwc
//	c.w = rwc
//	if debugServerConnections {
//		c.rwc = newLoggingConn("server", c.rwc)
//	}
//	c.sr = liveSwitchReader{r: c.rwc}
//	c.lr = io.LimitReader(&c.sr, noLimit).(*io.LimitedReader)
//	br := newBufioReader(c.lr)
//	bw := newBufioWriterSize(checkConnErrorWriter{c}, 4<<10)
//	c.buf = bufio.NewReadWriter(br, bw)
	return c, nil
}

func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		rw, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
//				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		c, err := srv.newConn(rw)
		if err != nil {
			continue
		}
//		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve()
	}
}

//----------------------------------------
func checkError(err error) {
	if err != nil {
		ERR("Fatal error:", err)
		os.Exit(-1)
	}
}