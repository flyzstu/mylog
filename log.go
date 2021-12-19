package mylog

// package main

import (
	"fmt"
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

// 日志通道
var ch chan string

// 结束通道
var fi chan struct{}

// 初始化通道
func init() {
	ch = make(chan string, 100) // 缓冲区大小为100的通道
	fi = make(chan struct{})    // 同步通道
	// 启动一个后台一直执行的goroutine，用来打印日志信息
	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
		fi <- struct{}{}
	}()
}

func (o *object) Close() {
	close(ch) // 关闭通道ch
	<-fi      // 阻塞主goroutine，直到通道内的数据发送完毕
}

// 获得当前时间
func getTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 新建一个日志对象
func New() *object {
	return &object{}
}

// 设置日志级别
func (o *object) SetLevel(level Level) {
	o.level = level
}

// 检查日志级别是否需要打印
func (o *object) checkLevel(level Level) bool {
	return o.level <= level
}

// 打印INFO级别日志
func (o *object) Info(msg string, a ...interface{}) {
	if o.checkLevel(INFO) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [INFO] %s", getTimeNow(), fmt.Sprintf(msg, a...))
		ch <- msg
	}
}

// 打印DEBUG级别日志
func (o *object) Debug(msg string, a ...interface{}) {
	if o.checkLevel(DEBUG) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [DEBUG] %s", getTimeNow(), fmt.Sprintf(msg, a...))
		ch <- msg
	}
}

// 打印WARNING级别日志
func (o *object) Warning(msg string, a ...interface{}) {
	if o.checkLevel(WARNING) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [WARNING] %s", getTimeNow(), fmt.Sprintf(msg, a...))
		ch <- msg
	}
}

// 打印ERROR级别日志
func (o *object) Error(msg string, a ...interface{}) {
	if o.checkLevel(ERROR) {
		// 拼接字符串
		msg = fmt.Sprintf("%s [ERROR] %s", getTimeNow(), fmt.Sprintf(msg, a...))
		ch <- msg
		ch <- msg
	}
}
