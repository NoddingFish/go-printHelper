上线打包：

```go
go build -i -tags tempdll -ldflags="-H windowsgui"
```



```
go build -i -tags tempdll              // dll 包一起打包
go build -i -ldflags="-H windowsgui"   // 去掉黑框(cmd窗口)
```

