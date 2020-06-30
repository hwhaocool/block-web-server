package main

import (
    "net"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"

    "fmt"
    "time"
    "io/ioutil"
    
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

    router.Any("/block:s", block )

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

    body,_ := ioutil.ReadAll(ctx.Request.Body)
    fmt.Println("---body/--- \r\n "+string(body))

    ctx.JSON(200, gin.H{
        "type":    "ok",
        "message": "web server is ok",
        "ip":      getLocalIP(),
    })
}

// block11 阻塞11秒的接口
func block11(ctx *gin.Context) {
    Logger.Info("block11",
        zap.String("remote addr", ctx.Request.RemoteAddr),
        zap.String("path", ctx.Request.URL.Path))

    body,_ := ioutil.ReadAll(ctx.Request.Body)
    fmt.Println("---body/--- \r\n "+string(body))
    
    time.Sleep(time.Duration(11)*time.Second)

    Logger.Info("block11 block end")

    ctx.JSON(200, gin.H{
        "type":    "ok",
        "message": "grey proxy is ok",
        "ip":      getLocalIP(),
    })
}

// block 阻塞x秒的接口
func block(ctx *gin.Context) {
    str := ctx.Param("s")
    if 0 == len(str) {
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

    body,_ := ioutil.ReadAll(ctx.Request.Body)
    fmt.Println("---body/--- \r\n "+string(body))
    
    time.Sleep(time.Duration(seconds)*time.Second)

    Logger.Info("block end")

    ctx.JSON(200, gin.H{
        "type":    "ok",
        "message": "block ok",
        "ip":      getLocalIP(),
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
