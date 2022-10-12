package services

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/aditya109/animated-eureka/internal/models"
	"github.com/aditya109/animated-eureka/internal/utils"
	logger "github.com/sirupsen/logrus"
)

func getVirtualBonds(count int) []models.AvailableVirtualBond {
	virtualBonds := []models.AvailableVirtualBond{}
	for i := 0; i < count; i++ {
		virtualBonds = append(virtualBonds, models.AvailableVirtualBond{
			VirtualBondId: utils.GenerateRandomID(true),
			VirtualBond:   utils.GenerateRandomID(false),
		})
	}
	return virtualBonds
}

func InsertBulkVirtualBonds(db *sql.DB, count int) error {
	var (
		inserts []string
		vals    []interface{}
	)
	query := "INSERT INTO available_virtual_bond(virtual_bond_id, virtual_bond) VALUES "

	virtual_bonds := getVirtualBonds(count)

	for _, v := range virtual_bonds {
		inserts = append(inserts, "(?, ?)")
		vals = append(vals, v.VirtualBondId, v.VirtualBond)
	}

	queryVals := strings.Join(inserts, ",")
	query = query + queryVals

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		logger.Errorf("Error %s when preparing SQL statement", err)
		return err
	}
	defer statement.Close()
	result, err := statement.ExecContext(ctx, vals...)
	if err != nil {
		logger.Errorf("Error %s when inserting row into available_virtual_bond table", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		logger.Errorf("Error %s when finding rows affected", err)
		return err
	}
	logger.Infof("%d virtual_bonds got bulk inserted", rows)
	return nil
}
