// 由res2go IDE插件自动生成。
package main

import (
    _ "fmt"
    "github.com/ying32/govcl/vcl"
    _ "runtime"
)

func main() {
    vcl.Application.SetScaled(true)
    vcl.Application.SetTitle("打印助手")
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&Form1)

    SetTrayIcon(Form1)

    vcl.Application.Run()
}
