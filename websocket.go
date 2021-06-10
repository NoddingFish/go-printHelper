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
	Nick    string
	SubNick string
}

var f *TForm1
var Nick string
var SubNick string

func webSocketConn(wg *sync.WaitGroup, msg []byte) {
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial("wss://test.huijiedan.cn/websocket?type=print", nil)

	if err != nil {
		fmt.Println(err)

		return
	}

	werr := conn.WriteMessage(websocket.TextMessage, msg)

	if werr != nil {
		fmt.Println(werr)
	}
	//申明定时器10s，设置心跳时间为10s
	//ticker := time.NewTicker(time.Second * 10)
	//
	//connect := &Connetion{
	//	con: conn,
	//}
	//开启多线程
	//go connect.timeWriter(ticker, conn)

	//开启多线程
	go func() {
		for {
			_, message, err := conn.ReadMessage()

			if err != nil {
				fmt.Printf("read err:%v \n", err)
				f.LogBox.Items().Add(DateStr() + "：断开链接")
				//TODO 保持 listBox 滚动条显示最新记录
				f.LogBox.SetTopIndex(f.LogBox.Count() - 1)
				
				_ = conn.Close()
				return
			}
			//str := fmt.Sprintf("%s", message)
			//f.LogBox.Items().Add(str)

			jsonData := ByteToMap(message)
			typeD := jsonData["type"]

			switch typeD {
			case "login":
				f.LogBox.Items().Add(DateStr() + "：恭喜你，登录成功！")
				break
			case "online":
				onlineData := make(map[string]interface{})

				onlineData["source_client"] = "printHelper"
				onlineData["type"] = "online_back"
				onlineData["name"] = Nick
				onlineData["sub_name"] = SubNick
				onlineData["status"] = "success"

				onlineMsg, _ := json.Marshal(onlineData)

				werr := conn.WriteMessage(websocket.TextMessage, onlineMsg)

				if werr != nil {
					fmt.Println(werr)
				}
				break
			}

			fmt.Printf("ReadMessage:%s \n", message)

			fmt.Println(jsonData)
			fmt.Println(jsonData["type"])

			//TODO 保持 listBox 滚动条显示最新记录
			f.LogBox.SetTopIndex(f.LogBox.Count() - 1)
		}
	}()

}

func DateStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ByteToMap(byteData []byte) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal(byteData, &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
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

	msg["source_client"] = "printHelper"
	msg["type"] = "login"
	msg["name"] = Nick
	msg["sub_name"] = SubNick

	bMsg, _ := json.Marshal(msg)

	return bMsg
}

func WebsocketRun(fIn *TForm1, nick string, subNick string) {

	f = fIn
	Nick = nick
	SubNick = subNick

	flag.Parse()           //命令行参数
	var wg *sync.WaitGroup //申明计数器

	webSocketConn(wg, NewConnMsg())

	//wg.Wait() //阻塞代码的运行，直到计数器地值减为0
}

//func main() {
//	//NewConnMsg()
//	WebsocketRun()
//}
