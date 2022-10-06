package requests

import "github.com/aditya109/animated-eureka/internal/models"

type PostBulkVirtualBondsRequestWrapper struct {
	AvailableVirtualBonds []models.AvailableVirtualBond `json:"available_virtual_bonds"`
}
