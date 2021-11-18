package log

import (
	"Amanisis/config"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)

type Log struct {
	LogFile *os.File
	LogDebugFile *os.File
}
var LogClient *Log

func LogInit(LogConf config.LogServer){
	LogClient = new(Log)
	LogClient.LogFile, _ = os.OpenFile(LogConf.LogUrl,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	LogClient.LogDebugFile, _ = os.OpenFile(LogConf.DebugLogUrl,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
}

func (*Log) Record(ctx *gin.Context){
	log.SetOutput(LogClient.LogFile)
	body, _:= ioutil.ReadAll(ctx.Request.Body)
	rawStr := string(body)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	log.Println(ctx.ClientIP(),ctx.FullPath(),rawStr)
}

func (*Log) Debug(ctx *gin.Context,v ...interface{}){
	log.SetOutput(LogClient.LogDebugFile)
	log.Println(ctx.ClientIP(),ctx.FullPath(),v)
}

func (*Log) Info(ctx *gin.Context,v ...interface{}){
	log.SetOutput(LogClient.LogFile)
	log.Println(ctx.ClientIP(),ctx.FullPath(),v)
}