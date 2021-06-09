package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type Connetion struct {
	con   *websocket.Conn
	mutex sync.Mutex
}

type ConnectDate struct {
	Nick string
	SubNick string
}

func webSocketConn(wg *sync.WaitGroup, msg []byte) {
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial("wss://test.huijiedan.cn/websocket?type=print", nil)

	if err != nil {
		fmt.Println(err)

		return
	}

	werr := conn.WriteMessage(websocket.TextMessage, msg)

	msg1 := make(map[string]interface{})
	msg1["type"] = "login"
	msg1["name"] = "wangweilon"
	msg1["sub_name"] = "wangweilon:徐然"

	aMsg, _ := json.Marshal(msg1)

	fmt.Printf("测试 %s\n", aMsg)

	if werr != nil {
		fmt.Println(werr)
	}
	//申明定时器10s，设置心跳时间为10s
	ticker := time.NewTicker(time.Second * 10)

	connect := &Connetion{
		con: conn,
	}
	//开启多线程
	go connect.timeWriter(ticker, conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("read err:%v \n", err)
			return
		}
		fmt.Printf("read:%s \n", message)

		//互斥锁
		connect.mutex.Lock()
		werr2 := connect.con.WriteMessage(websocket.TextMessage, aMsg)
		connect.mutex.Unlock()

		if werr2 != nil {
			fmt.Println(werr2)
		}

		fmt.Printf("received: %s\n", message)
	}
	wg.Done() // 每次把计数器-1

}

func (con *Connetion) timeWriter(ticker *time.Ticker, c *websocket.Conn) {

	for {
		<-ticker.C
		err := c.SetWriteDeadline(time.Now().Add(10 * time.Second))
		//fmt.Println(time.Now().Format(time.UnixDate))
		if err != nil {
			log.Printf("ping error: %s\n", err.Error())
		}

		con.mutex.Lock()
		if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
			log.Printf("ping error: %s\n", err.Error())
		}
		con.mutex.Unlock()

	}
}

func NewConnMsg() []byte {

	msg := make(map[string]interface{})


	msg["type"] = "login"
	msg["name"] = "wangweilon"
	msg["sub_name"] = "wangweilon:徐然"

	bMsg, _ := json.Marshal(msg)

	return bMsg
}

func WebsocketRun() {

	flag.Parse()           //命令行参数
	var wg *sync.WaitGroup //申明计数器

	webSocketConn(wg, NewConnMsg())

	wg.Wait() //阻塞代码的运行，直到计数器地值减为0
}

func main() {
	//NewConnMsg()
	WebsocketRun()
}
