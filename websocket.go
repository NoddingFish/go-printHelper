package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ying32/govcl/vcl"
	_ "log"
	"time"
)

type Connect struct {
	conn *websocket.Conn
}

var f *TForm1
var Nick string
var SubNick string
var CNConnect *websocket.Conn
var WebSocketConnect *websocket.Conn

type PrintData struct {
	A uint32 `json:"a"`
	B string `json:"b"`
	C uint32 `json:"c"`
}

func CNPrintConnect() *websocket.Conn {
	//TODO 连接菜鸟打印组件
	var dialer *websocket.Dialer
	CNconn, _, err := dialer.Dial("wss://localhost:13529", nil)

	if err != nil {
		vcl.ShowMessage("未连接菜鸟打印组件，请启动菜鸟打印组件！")
		f.Button1.SetEnabled(true)
		f.Button2.SetEnabled(false)
		return nil
	}

	return CNconn
}

func webSocketConnect() *websocket.Conn {
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial("wss://test.huijiedan.cn/websocket?type=print", nil)

	if err != nil {
		vcl.ShowMessage("连接打印服务失败！")
		f.Button1.SetEnabled(true)
		f.Button2.SetEnabled(false)
		return nil
	}
	return conn
}

func webSocketListen(msg []byte) {

	werr := WebSocketConnect.WriteMessage(websocket.TextMessage, msg)

	if werr != nil {
		fmt.Println(werr)
	}

	//开启多线程
	go func() {
		for {
			_, message, err := WebSocketConnect.ReadMessage()

			if err != nil {
				fmt.Printf("read err:%v \n", err)
				f.LogBox.Items().Add(DateStr() + "：断开链接")
				//TODO 保持 listBox 滚动条显示最新记录
				f.LogBox.SetTopIndex(f.LogBox.Count() - 1)
				f.Button1.SetEnabled(true)
				f.Button2.SetEnabled(false)
				_ = WebSocketConnect.Close()
				return
			}
			//str := fmt.Sprintf("%s", message)
			//f.LogBox.Items().Add(str)

			jsonData := ByteToMap(message)
			typeD := jsonData["type"]

			switch typeD {
			case "login":
				f.LogBox.Items().Add(DateStr() + "：恭喜你，登录成功！")
			case "online":
				onlineMsg := NewConnMsg("online_back")
				werr := WebSocketConnect.WriteMessage(websocket.TextMessage, onlineMsg)

				if werr != nil {
					fmt.Println(werr)
				}
			case "print":
				//	jsonData["cmd"] = "print"
				//	jsonData["requestID"] = "1231"
				//	jsonData["version"] = "1.0"
				//	jsonData["task"] = {
				//		"taskID": "12313",
				//		"preview": "12313",
				//		"printer": "12313",
				//		"notifyType": ["print"],
				//		"previewType": "image",
				//		"documents": jsonData
				//}

				//var PrintData map[string]interface{}

				//err1 := json.Unmarshal(message, &PrintData)
				//if err1 != nil {
				//	fmt.Println(err1)
				//}
				//PrintData["cmd"] = "print"
				//PrintData["requestID"] = "123123"
				//PrintData["version"] = "1.0"
				//PrintData["task"] = map[string]interface{}
				//
				//output, err := json.Marshal(obj)

				PrintData, _ := json.Marshal(jsonData)
				cerr := WebSocketConnect.WriteMessage(websocket.TextMessage, PrintData)
				if cerr != nil {
					fmt.Println(cerr)
				}

				onlineMsg := NewConnMsg("print_back")

				werr := WebSocketConnect.WriteMessage(websocket.TextMessage, onlineMsg)

				if werr != nil {
					fmt.Println(werr)
				}
			}

			fmt.Printf("ReadMessage:%s \n", message)

			fmt.Println(jsonData)
			fmt.Println(jsonData["type"])

			//TODO 保持 listBox 滚动条显示最新记录
			f.LogBox.SetTopIndex(f.LogBox.Count() - 1)
		}
	}()

	//心跳
	go func() {
		for {
			werr := WebSocketConnect.WriteMessage(websocket.TextMessage, NewConnMsg("heart"))

			if werr != nil {
				fmt.Println(werr)
			}
			fmt.Println("发送心跳~")
			time.Sleep(time.Second * 20)
		}
	}()

}

// 时间格式字符串
func DateStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// byte to map
func ByteToMap(byteData []byte) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal(byteData, &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}

//func (con *Connetion) timeWriter(ticker *time.Ticker, c *websocket.Conn) {
//
//	for {
//		<-ticker.C
//		err := c.SetWriteDeadline(time.Now().Add(10 * time.Second))
//		//fmt.Println(time.Now().Format(time.UnixDate))
//		if err != nil {
//			log.Printf("ping error: %s\n", err.Error())
//		}
//
//		con.mutex.Lock()
//		if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
//			log.Printf("ping error: %s\n", err.Error())
//		}
//		con.mutex.Unlock()
//
//	}
//}

//{"source_client":"printHelper","type":"login","name":"wangweilon","sub_name":"wangweilon:徐然"}

//{"source_client":"printHelper","type":"heart"}

//{"source_client":"printHelper","type":"online","name":"wangweilon","sub_name":"wangweilon:徐然"}

//{"source_client":"printHelper","type":"online_back","name":"wangweilon","sub_name":"wangweilon:徐然","status":"success"}

//{"documentID":"1622795822211001","contents":[{"templateURL":"http:\/\/cloudprint.cainiao.com\/template\/standard\/372806\/3","data":{"recipient":{"address":{"province":"河南省","city":"南阳市","district":"卧龙区","town":"","detail":"七里园乡永和苑B10栋6楼"},"name":"蔡博","phone":"17681843982"},"routingInfo":{"consolidation":{"name":"集包地：南阳公司包  ","code":"KR-CN"},"sortation":{"name":""},"routeCode":"721-A010-J9 47"},"sender":{"address":{"province":"安徽省","city":"安庆市","district":"大观区","detail":"菱湖南路238号"},"name":"啊龙","mobile":"13865188214","phone":""},"trade_id":"1847946854378423439","waybillCode":"7790003925781","repeat":0,"item":{"order_no":"1847946854378423439","port_img":"http:\/\/cdn-cloudprint.cainiao.com\/waybill-print\/cloudprint-imgs\/6a39502c7c0d48ba96f3f6340988648b.jpg","products":[{"name":"11×1"}],"desc":"测试\n"},"is_vip":"0"}},{"templateURL":"http:\/\/cloudprint.cainiao.com\/template\/standard\/330032\/3","data":{"products":[{"name":"11×1"}],"port_img":"http:\/\/cdn-cloudprint.cainiao.com\/waybill-print\/cloudprint-imgs\/6a39502c7c0d48ba96f3f6340988648b.jpg"}}],"is_websocket":true,"source_client":"printHelper","type":"print","printer":"Fax","source":"SJZS","name":"wangweilon","sub_name":"wangweilon:dkj","taobao_user_id":"35255526"}

//{"source_client":"printHelper","order_no":"1622795944211001","status":"success","type":"print_back"}
func NewConnMsg(typeStr string) []byte {

	msg := make(map[string]interface{})

	msg["source_client"] = "printHelper"
	msg["type"] = typeStr
	msg["name"] = Nick
	msg["sub_name"] = SubNick

	switch typeStr {
	case "heart", "login":
		break
	case "online_back", "print_back":
		msg["status"] = "success"
		break
	}

	bMsg, _ := json.Marshal(msg)

	fmt.Printf("发送消息:%s \n", bMsg)
	return bMsg
}

func WebsocketRun(fIn *TForm1, nick string, subNick string) {

	f = fIn
	Nick = nick
	SubNick = subNick

	CNConnect = CNPrintConnect() // 连接菜鸟打印组件
	if CNConnect != nil {
		WebSocketConnect = webSocketConnect() // 连接 websocket 服务
		if CNConnect != nil {
			webSocketListen(NewConnMsg("login"))
		}
	}
}

//func main() {
//	//NewConnMsg()
//	WebsocketRun()
//}
