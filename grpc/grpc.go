package grpc

import (
	"github.com/arthasyou/net-go/grpc"
	"github.com/spf13/viper"
)

// Connect to grpc server
func Connect() {
	consulAddr := viper.GetString("Consul.addr")
	serviceName := viper.GetString("Service.name")
	grpc.Connect(consulAddr, serviceName)
}
