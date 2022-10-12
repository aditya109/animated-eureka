package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	logger "github.com/sirupsen/logrus"
)

var DB *sql.DB

func EstablishConnectionToDatabase(driver string, connectionString string) error {
	var err error
	DB, err = sql.Open(driver, connectionString)
	if err != nil {
		logger.Error("error while establishing connection to db", err)
		return err
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(time.Minute * 5)
	logger.Info("connection to database established.")
	return nil
}
