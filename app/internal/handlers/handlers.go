package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aditya109/animated-eureka/internal/models"
	"github.com/aditya109/animated-eureka/internal/responses"
	resp "github.com/aditya109/animated-eureka/internal/responses"
	svc "github.com/aditya109/animated-eureka/internal/services"
	logger "github.com/sirupsen/logrus"
)

// WelcomeHandler returns welcome message for home URL
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	responseStatusCode := 200
	var response resp.WelcomeResponseWrapper = "OK"
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(response))
	logger.Info(fmt.Sprintf("STATUS: %d === / route was hit", responseStatusCode))
}

// GetItemsHandler returns a static list of items, no query params required
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	var responseStatusCode int
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	items, err := svc.GetAllItems()
	if err != nil {
		responseStatusCode = http.StatusInternalServerError
		w.WriteHeader(responseStatusCode)
		logger.Error(err)
		w.Write([]byte(fmt.Sprintf("error occurred while getting items, %s", err.Error())))
	} else {
		responseStatusCode = http.StatusOK
		w.WriteHeader(responseStatusCode)
		json.NewEncoder(w).Encode(items)
	}
	logger.Info(fmt.Sprintf("STATUS: %d === /items route was hit", responseStatusCode))
}

// PostVirtualBondHandler returns all the details of the virtual bond created, provided a valid request body is provided.
func PostVirtualBondHandler(w http.ResponseWriter, r *http.Request) {
	var m responses.PostGetVidResponseWrapper = resp.PostGetVidResponseWrapper{
		VirtualBond: models.VirtualBond{
			VirtualBondId: "028dd2df5592459f8a5700bb0ce809ef",
			UID:           "32c335b4-ca07-4b12-8370-720a488f9ecc",
			AFactionIds: []string{
				"a5eb5d62dc984c2d82a375682eb8401e",
				"720c8d242d8f4566a6fc5479dc530b2d",
			},
			BFactionIds: []string{
				"e2a14e2d2af44133ba8ad41e3f286bce",
				"9847d5ab91dc4816968a8fa9be47608f",
				"bd0b81ddba064014813042048a53b82e",
			},
			BondEndTime: "4d",
		},
	}
	var responseStatusCode int
	w.Header().Set("Content-Type", "application/json")

	responseStatusCode = http.StatusOK
	w.WriteHeader(responseStatusCode)
	json.NewEncoder(w).Encode(m)
	logger.Info(fmt.Sprintf("STATUS: %d === /items route was hit", responseStatusCode))
}
