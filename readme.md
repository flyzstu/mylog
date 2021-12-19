### 一个基于go语言简易异步日志库

### 安装

```bash
go get github.com/flyzstu/mylog
```

### 使用

日志级别：DEBUG > INFO > WARNING > ERROR

```go
var logger = mylog.New()
logger.SetLevel(mylog.INFO) // 日志级别大于INFO的都可以显示
logger.Info("111")
logger.Debug("222") // 不显示
logger.Info("111")
logger.Warning("111") // 显示
```

