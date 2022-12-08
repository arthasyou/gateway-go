package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpc "github.com/arthasyou/gateway-go/grpc"
	"github.com/arthasyou/gateway-go/http"
	"github.com/arthasyou/gateway-go/nsq"
	"github.com/arthasyou/gateway-go/socket"
	"github.com/arthasyou/utility-go/logger"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("cfg")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: ", err)
	}
}

func initLog() {
	path := viper.GetString("Log.path")
	level := viper.GetString("Log.level")
	logger.InitLog(path, level)
}

func start() {
	grpc.Connect()
	nsq.Connect()
	socket.Start()
	http.Start()
}

func waitExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
}

func main() {
	initConfig()
	initLog()
	start()
	waitExit()
}
