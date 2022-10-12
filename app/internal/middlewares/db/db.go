package db

import (
	"database/sql"

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
	logger.Info("connection to database established.")
	return nil
}
