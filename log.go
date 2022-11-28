package mylog

// package main

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// 自定义log库

// 定义日志级别
type Level int

const (
	DEBUG Level = iota //常量计数器
	INFO
	WARNING
	ERROR
	NONE
)

// 日志对象
type object struct {
	level Level
}

var (
	ch               chan string   // 日志通道
	fi               chan struct{} // 结束信号通道
	defaultLogObject *object       // 默认日志对象
)

func logHelper() string {
	_, file, line, _ := runtime.Caller(2)
	file = path.Base(file)
	return fmt.Sprintf("[%v: %v]", file, line)
}

// 初始化通道
func init() {
	ch = make(chan string, 100) // 缓冲区大小为100的通道
	fi = make(chan struct{})    // 同步通道
	defaultLogObject = new(object)
	// 启动一个后台一直执行的goroutine，用来打印日志信息
	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
		fi <- struct{}{}
	}()
}

func Close() {
	close(ch) // 关闭通道ch
	<-fi      // 阻塞主goroutine，直到通道内的数据发送完毕
}

// 获得当前时间
func getTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 设置日志级别
func SetLevel(level Level) {
	defaultLogObject.level = level
}

// 检查日志级别是否需要打印
func checkLevel(level Level) bool {
	return defaultLogObject.level <= level
}

// 打印INFO级别日志
func Info(msg string, a ...interface{}) {
	if checkLevel(INFO) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [INFO] %v %v", getTimeNow(), logHelper(), fmt.Sprintf(msg, a...))
		ch <- msg
	}
}

// 打印DEBUG级别日志
func Debug(msg string, a ...interface{}) {
	if checkLevel(DEBUG) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [DEBUG] %v %v", getTimeNow(), logHelper(), fmt.Sprintf(msg, a...))
		ch <- msg
	}
}

// 打印WARNING级别日志
func Warning(msg string, a ...interface{}) {
	if checkLevel(WARNING) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [WARNING] %v %v", getTimeNow(), logHelper(), fmt.Sprintf(msg, a...))
		ch <- msg
	}
}

// 打印ERROR级别日志
func Error(msg string, a ...interface{}) {
	if checkLevel(ERROR) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [ERROR] %v %v", getTimeNow(), logHelper(), fmt.Sprintf(msg, a...))
		ch <- msg
	}
}
