package responses

import "github.com/aditya109/animated-eureka/internal/models"

// PostGetVidResponseWrapper contains all the details of the virtualbond returned in the response
// swagger:response PostGetVidResponse
type PostGetVidResponseWrapper struct {
	// All the details for virtual bond id
	// in: virtualbond
	VirtualBond models.VirtualBond
}
