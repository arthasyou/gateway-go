package grpc

import (
	"github.com/luobin998877/go_net/grpc"
	"github.com/spf13/viper"
)

// Connect to grpc server
func Connect() {
	consulAddr := viper.GetString("Consul.addr")
	serviceName := viper.GetString("Service.name")
	grpc.Connect(consulAddr, serviceName)
}
