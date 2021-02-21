package http

import (
	"github.com/arthasyou/net-go/http"
	"github.com/spf13/viper"
)

// Start http server
func Start() {
	port := viper.GetInt("Http.port")
	isSsl := viper.GetBool("Http.ssl")
	if isSsl {
		cert := viper.GetString("Ssl.CertFile")
		key := viper.GetString("Ssl.KeyFile")
		http.StartSsl(uint16(port), cert, key)
	} else {
		http.Start(uint16(port))
	}
}
