package main

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"fmt"
	"io/ioutil"
	"time"

	"strconv"
)

func main() {
	//初始化日志
	InitLogger()
	Logger.Info("simple web server begin to start")

	//打印ip
	Logger.Info("local ip", zap.String("ip", getLocalIP()))

	//新建gin 实例
	router := gin.New()

	router.Any("/block:s", block)
	router.GET("/kb/:s", KB)
	router.GET("/mb/:s", MB)

	//其它 -> 进行分发
	router.NoRoute(welcome)

	//启动 gin 并监听端口
	err := router.Run(":8080")
	if err != nil {
		Logger.Fatal("proxy start failed,", zap.Error(err))
	}
}

// welcome 健康检查接口
func welcome(ctx *gin.Context) {
	Logger.Info("welcome",
		zap.String("remote addr", ctx.Request.RemoteAddr),
		zap.String("path", ctx.Request.URL.Path))

	body, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))

	ctx.JSON(200, gin.H{
		"type":    "ok",
		"message": "web server is ok",
		"ip":      getLocalIP(),
	})
}

// block 阻塞x秒的接口
func block(ctx *gin.Context) {
	str := ctx.Param("s")
	if str == "" {
		Logger.Warn("block  time is empty")
		welcome(ctx)
		return
	}

	seconds, err := strconv.Atoi(str)

	if err != nil {
		welcome(ctx)
		return
	}

	Logger.Info("block",
		zap.Int("seconds", seconds),
		zap.String("remote addr", ctx.Request.RemoteAddr),
		zap.String("path", ctx.Request.URL.Path))

	body, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))

	time.Sleep(time.Duration(seconds) * time.Second)

	Logger.Info("block end")

	ctx.JSON(200, gin.H{
		"type":    "ok",
		"message": "block ok",
		"ip":      getLocalIP(),
	})
}

// 返回 x KB的响应体
func KB(ctx *gin.Context) {
	str := ctx.Param("s")
	if str == "" {
		Logger.Warn("kb size is empty")
		welcome(ctx)
		return
	}

	size, err := strconv.Atoi(str)

	if err != nil {
		welcome(ctx)
		return
	}

	Logger.Info("KB",
		zap.Int("size", size),
		zap.String("remote addr", ctx.Request.RemoteAddr),
		zap.String("path", ctx.Request.URL.Path))

	ctx.JSON(200, gin.H{
		"code":    0,
		"message": strings.Repeat("*", size*1024),
	})
}

// 返回 x MB的响应体
func MB(ctx *gin.Context) {
	str := ctx.Param("s")
	if str == "" {
		Logger.Warn("MB size is empty")
		welcome(ctx)
		return
	}

	size, err := strconv.Atoi(str)

	if err != nil {
		welcome(ctx)
		return
	}

	Logger.Info("MB",
		zap.Int("size", size),
		zap.String("remote addr", ctx.Request.RemoteAddr),
		zap.String("path", ctx.Request.URL.Path))

	ctx.JSON(200, gin.H{
		"code":    0,
		"message": strings.Repeat("*", size*1024*1024),
	})
}

// getLocalIP 得到local ip
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
