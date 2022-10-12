package server

import (
	"fmt"
	"time"

	"github.com/aditya109/animated-eureka/internal/models"
)

var (
	configuration models.ServerConfig
	httpPort      string
	prefix        string
	endpoint      string
	writeTimeout  time.Duration
	readTimeout   time.Duration
)

func InitializeServerConfiguration(config *models.ServerConfig) {
	setHTTPPortFromConfigObject() // getting http port from config
	setEndpointFromConfigObject() // getting endpoint from config
	setTimeoutsFromConfigObject() // getting timeouts from config
}

// setHTTPPortFromConfigObject sets httpPort variable to port mentioned in the config object
func setHTTPPortFromConfigObject() {
	httpPort = configuration.ServerPort
}

// setEndpointFromConfigObject sets endpoint IP from config object
func setEndpointFromConfigObject() {
	if configuration.IsTLSEnabled {
		prefix = "https"
	} else {
		prefix = "http"
	}
	endpoint = fmt.Sprintf("%s:%s", "localhost", httpPort)
}

// setTimeoutsFromConfigObject sets writeTimeout and readTimeout from config object
func setTimeoutsFromConfigObject() {
	writeTimeout = time.Duration(configuration.WriteTimeout) * time.Second
	readTimeout = time.Duration(configuration.ReadTimeout) * time.Second
}
