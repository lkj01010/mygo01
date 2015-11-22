package main

import (


	"golang.org/x/net/websocket"
	"time"
	"math/rand"
	"net/http"
	"strconv"
	"encoding/json"
	"os"
	"mygo01/cfg"
)

var Config cfg.Configuration

func load(configfile string) cfg.Configuration {
	config := cfg.Configuration{}
	file, _ := os.Open(configfile)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func wsHandler(ws *websocket.Conn) {
//	//p := unsafe.Pointer(&ws)
//	//index := ((int)(uintptr(p))) % n
//	index := rand.Intn(n)
//	lock := locks[index]
//	lock.Lock()
//	wsList[index].PushBack(ws)
//	lock.Unlock()
//
//	for {
//		var reply string
//		if err := websocket.Message.Receive(ws, &reply); err != nil {
//			fmt.Println("Can't receive because of " + err.Error())
//			break
//		}
//	}
//
//	lock.Lock()
//	for e := wsList[index].Front(); e != nil; e = e.Next() {
//		if e.Value.(*websocket.Conn) == ws {
//			wsList[index].Remove(e)
//			break
//		}
//	}
//	lock.Unlock()
}

func main() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	Config = load("config.json")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(wsHandler)}
		s.ServeHTTP(w, req)
	})

	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
	if err != nil {
		panic("Error: " + err.Error())
	}

}


