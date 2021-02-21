package socket

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/arthasyou/net-go/packet"
	"github.com/arthasyou/net-go/socket"
	"github.com/spf13/viper"
)

// Start all socket
func Start() {
	reigsterNode()
	packet.Register(&handler{})
	loadCfg()
}

func reigsterNode() {
	node := viper.GetString("Node.name")
	socket.RegisterNode(node)
}

func loadCfg() {
	m := viper.Get("Socket")
	switch reflect.TypeOf(m).Kind() {

	case reflect.Slice:
		s := reflect.ValueOf(m)
		for i := 0; i < s.Len(); i++ {
			so := s.Index(i)
			si := so.Interface()
			sm := si.(map[string]interface{})
			port := readPort(sm["port"])
			genre := readGenre(sm["genre"])
			isSsl := readSsl(sm["ssl"])
			startSocketServer(genre, port, isSsl)
		}
	}
}

func readPort(port interface{}) uint16 {
	s := fmt.Sprintf("%v", port)
	i, _ := strconv.ParseUint(s, 10, 16)
	return uint16(i)
}

func readGenre(genre interface{}) string {
	s := fmt.Sprintf("%v", genre)
	return s
}

func readSsl(ssl interface{}) bool {
	s := fmt.Sprintf("%v", ssl)
	b, _ := strconv.ParseBool(s)
	return b
}

func startSocketServer(genre string, port uint16, isSsl bool) {
	switch genre {
	case "tcp":
		go socket.StartTCP(port)
	case "ws":
		if isSsl {
			cert := viper.GetString("Ssl.CertFile")
			key := viper.GetString("Ssl.KeyFile")
			socket.StartWss(port, cert, key)
		} else {
			socket.StartWs(port)
		}
	}
}
