package main

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/aditya109/animated-eureka/internal/server"
	cfg "github.com/aditya109/animated-eureka/pkg/config"
	logCfg "github.com/aditya109/animated-eureka/pkg/logger"
	logger "github.com/sirupsen/logrus"
)

func getProjectName() string {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
		separator  = ""
	)
	if strings.Contains(basepath, "\\") {
		separator = "\\"
	} else {
		separator = "/"
	}
	tree := strings.Split(basepath, separator)
	return tree[len(tree)-1]
}

func main() {
	var PROJECT_NAME = getProjectName()
	config, err := cfg.GetConfiguration(PROJECT_NAME) // retrieving configuration
	if err != nil {
		logger.Fatal(err)
	}
	logCfg.InitializeLogging(config) // initializing logger
	server.Start(config)

}
