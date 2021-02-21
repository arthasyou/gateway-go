package nsq

import "fmt"

type consumerHandler struct{}

func (h *consumerHandler) ProcessMessage(topic string, channel string, message []byte) error {
	fmt.Println("topic: ", topic, "channel: ", channel, "message: ", message)
	return nil
}
