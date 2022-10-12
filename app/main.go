package main

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/aditya109/animated-eureka/internal/middlewares/db"
	"github.com/aditya109/animated-eureka/internal/server"
	cfg "github.com/aditya109/animated-eureka/pkg/config"
	"github.com/aditya109/animated-eureka/pkg/helper"
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

	// starting database connection
	databaseConnectionString := helper.GetFormattedConnectionStringFromDatabaseConfig(config.DatabaseConfig)
	err = db.EstablishConnectionToDatabase(config.DatabaseConfig.DriverName, databaseConnectionString)
	if err != nil {
		log.Fatal("Error in connecting to database")
	}

	server.InitializeServerConfiguration(&config.ServerConfig)
	server.Start()
}
