package models

import "github.com/gofrs/uuid"

// GenerateUuid is used for creating unique IDs to be used mainly in generated HTML elements
func GenerateUuid() (string, error) {
	curId, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return curId.String(), nil
}
