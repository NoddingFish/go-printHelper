// 由res2go IDE插件自动生成。
package main

import (
	"fmt"
	_ "fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/win"
	"github.com/ying32/govcl/vcl/win/errcode"
	"os"
	_ "runtime"
	"syscall"
	"unsafe"
)

var (
	kernel32dll  = syscall.NewLazyDLL("kernel32.dll")
	_CreateMutex = kernel32dll.NewProc("CreateMutexW")
)

// 不知道为什么GetLastError无法获取，只能重新申明下
func CreateMutex(lpMutexAttributes *win.TSecurityAttributes, bInitialOwner bool, lpName string) (uintptr, uintptr, error) {
	return _CreateMutex.Call(uintptr(unsafe.Pointer(lpMutexAttributes)), win.CBool(bInitialOwner), win.CStr(lpName))
}

func main() {
	Mutex, _, err := CreateMutex(nil, true, "打印助手")
	defer win.ReleaseMutex(Mutex)
	fmt.Println("Mutex:", Mutex, err)
	if errNo, ok := err.(syscall.Errno); ok && errNo == errcode.ERROR_ALREADY_EXISTS {
		win.MessageBox(0, "您已经运行一个打印助手啦！可在右下角查看。", "重复运行", win.MB_OK+win.MB_ICONINFORMATION)
		os.Exit(1)
	}

	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("打印助手")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&Form1)

	SetTrayIcon(Form1)

	vcl.Application.Run()
}
