package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpc "github.com/luobin998877/go_gateway/gRPC"
	"github.com/luobin998877/go_gateway/http"
	"github.com/luobin998877/go_gateway/nsq"
	"github.com/luobin998877/go_gateway/socket"
	"github.com/luobin998877/go_utility/logger"
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
