package utils

import (
	"strings"

	"github.com/google/uuid"
)

// UUID is represented as 32 hexadecimal digits(Base 16) displayed in 5 parts separated by hyphens. The form is 8-4-4-4-12. So there are a total of 36 characters which include 32 hexadecimal and 4 hyphens.
func GenerateRandomID(isObjectId bool) string {
	uuidWithHyphen := uuid.New().String()
	if isObjectId {
		return uuidWithHyphen
	} else {
		return strings.Replace(uuidWithHyphen, "-", "", -1)
	}
}
