package main

import (
	"github.com/ying32/govcl/vcl"
	"time"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {

}

func (f *TForm1) OnLabel1Click(sender vcl.IObject) {

}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	nick := f.Nick.Text()
	SubNick := f.SubNick.Text()
	if len(nick) == 0 {
		vcl.ShowMessage("主账号必填！")
		return
	}
	if len(SubNick) == 0 {
		vcl.ShowMessage("子账号必填！")
		return
	}
	f.Button1.SetEnabled(false)
	f.Button2.SetEnabled(true)
	f.LogBox.Items().Add(time.Now().Format("2006-01-02 15:04:05") + "：提交数据！")
	WebsocketRun(f,nick, SubNick)
}

func (f *TForm1) OnLogClearClick(sender vcl.IObject) {
	f.LogBox.Items().Clear()//清空日志
}
