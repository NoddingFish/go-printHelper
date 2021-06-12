package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"runtime"
	"time"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	// 关闭窗口的询问事件
	f.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		//f.Hide()
		//f.SetWindowState(types.WsMinimized)
		//*canClose = false
		if vcl.MessageDlg("是否最小化到托盘？", types.MtConfirmation, types.MbYes, types.MbClose) == types.IdYes {
			f.Hide()
			*canClose = false
		}
	})
}

func SetTrayIcon(f *TForm1) {

	trayicon := vcl.NewTrayIcon(f)

	pm := vcl.NewPopupMenu(f)
	item := vcl.NewMenuItem(f)
	item.SetCaption("显示")
	item.SetOnClick(func(vcl.IObject) {
		f.Show()
	})
	pm.Items().Add(item)

	item2 := vcl.NewMenuItem(f)
	item2.SetCaption("退出")
	item2.SetOnClick(func(vcl.IObject) {
		f.Close()

	})
	pm.Items().Add(item2)
	trayicon.SetPopupMenu(pm)

	trayicon.SetHint(f.Caption())
	trayicon.SetVisible(true)

	//// 捕捉最小化
	//vcl.Application.SetOnMinimize(func(sender vcl.IObject) {
	//	f.Hide() // 主窗口最隐藏掉
	//})

	// 这里写啥好呢，macOS下似乎这些事件跟PopupMenu有冲突
	if runtime.GOOS != "darwin" {
		trayicon.SetOnDblClick(func(vcl.IObject) {
			// macOS似乎不支持双击
			//trayicon.SetBalloonTitle("打印通知")
			//trayicon.SetBalloonTimeout(20000)
			//trayicon.SetBalloonHint("我是提示正文啦")
			//trayicon.ShowBalloonHint()
			//fmt.Println("TrayIcon DClick.")
			f.Show()
		})
	}

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
	WebsocketRun(f, nick, SubNick)
}

func (f *TForm1) OnLogClearClick(sender vcl.IObject) {
	f.LogBox.Items().Clear() //清空日志
}
