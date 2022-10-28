package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// GenerateUuid is used for creating unique IDs to be used mainly in generated HTML elements
func GenerateUuid() (string, error) {
	curId, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return curId.String(), nil
}

// ConvertToDateTime inputs a string formatted as "YYYY-MM-DDTHH:mm" and returns a parsed time.Time value
func ConvertToDateTime(dateString string) (time.Time, error) {
	layout := "2006-01-02T15:04"
	return time.Parse(layout, dateString)
}

// SplitTasksByStatus inputs a slice of Task and returns a slice of Task for each status type
func SplitTasksByStatus(tasks []*Task) ([]*Task, []*Task, []*Task) {
	var ns, ip, com []*Task
	for _, t := range tasks {
		switch t.Status {
		case NOTSTARTED:
			ns = append(ns, t)
		case INPROGRESS:
			ns = append(ip, t)
		case COMPLETED:
			ns = append(com, t)
		}
	}
	return ns, ip, com
}
