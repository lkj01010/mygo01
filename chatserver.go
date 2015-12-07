package main

import (


//	"golang.org/x/net/websocket"
	"time"
//	"math/rand"
//	"net/http"
	"strconv"
//	"encoding/json"
//	"os"
//	"mygo01/cfg"
	"fmt"
	"gopkg.in/redis.v3"
)

//var Config cfg.Configuration
//
//func load(configfile string) cfg.Configuration {
//	config := cfg.Configuration{}
//	file, _ := os.Open(configfile)
//	decoder := json.NewDecoder(file)
//	err := decoder.Decode(&config)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return config
//}
//
//func wsHandler(ws *websocket.Conn) {
////	//p := unsafe.Pointer(&ws)
////	//index := ((int)(uintptr(p))) % n
////	index := rand.Intn(10)
////	lock := locks[index]
////	lock.Lock()
////	wsList[index].PushBack(ws)
////	lock.Unlock()
////
////	for {
////		var reply string
////		if err := websocket.Message.Receive(ws, &reply); err != nil {
////			fmt.Println("Can't receive because of " + err.Error())
////			break
////		}
////		fmt.Println(reply)
////	}
//
//	for{
//		msg := make([]byte, 12)
//		n, err := ws.Read(msg)
//		if err != nil {
//			fmt.Println("Can't read because of " + err.Error())
//		}
//		fmt.Printf("Receive: %s\n", msg[:n])
//
//		send_msg := "[" + string(msg[:n]) + "]"
//		_, err = ws.Write([]byte(send_msg))
//		if err != nil {
//			fmt.Println("Can't write because of " + err.Error())
//		}
//		fmt.Printf("Send: %s\n", send_msg)
//	}
//
//
////
////	lock.Lock()
////	for e := wsList[index].Front(); e != nil; e = e.Next() {
////		if e.Value.(*websocket.Conn) == ws {
////			wsList[index].Remove(e)
////			break
////		}
////	}
////	lock.Unlock()
//	fmt.Println("123")
//}
//
//func main0() {
//	seed := time.Now().UTC().UnixNano()
//	rand.Seed(seed)
//
//	Config = load("config.json")
//
//	fmt.Print("123")
//
//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//		s := websocket.Server{Handler: websocket.Handler(wsHandler)}
//		s.ServeHTTP(w, req)
//	})
//
//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
//	if err != nil {
//		panic("Error: " + err.Error())
//	}
//
//}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "203.195.243.33:12308",
//		Addr:     "42.62.101.24:12308",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	t1:=time.Now()
	for i:=0; i<100; i++{
		client.Set("root"+strconv.Itoa(i), strconv.Itoa(i), 0)
		fmt.Println(client.Get("root"+strconv.Itoa(i+1)))
	}
	t2:=time.Now()
	delta:=t2.Sub(t1)

	fmt.Println(delta)
}

