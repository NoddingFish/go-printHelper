
package main

import (
    "github.com/ying32/govcl/vcl"
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

    WebsocketRun(nick, SubNick)

    f.LogBox.Items().Add("恭喜你，登录成功！")
}

