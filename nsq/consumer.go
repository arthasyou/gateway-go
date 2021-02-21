package nsq

import (
	"github.com/arthasyou/utility-go/nsq"
	"github.com/spf13/viper"
)

// Connect nsq node
func Connect() {
	nsq.RegisterConsumerHandler(&consumerHandler{})
	addr := viper.GetString("Consumer.addr")
	topic := viper.GetString("Node.name")
	channel := viper.GetString("Consumer.channel")
	nsq.ConsumerStart(addr, topic, channel)
}

// Stop Cåonsumer
func Stop() {
	nsq.ConsumerStop()
}
