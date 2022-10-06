package responses

// PostBulkVirtualBondsResponseWrapper contains all the details of the virtualbond returned in the response
// swagger:response PostBulkVirtualBondsResponse
type PostBulkVirtualBondsResponseWrapper struct {
	// All the details for virtual bond id
	// in: virtualbond
	Status string
	Error  string
}
