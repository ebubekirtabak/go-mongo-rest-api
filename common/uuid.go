package common

import "github.com/google/uuid"

func GetUuid() (string, error) {
	uuidData, err := uuid.NewUUID()
	uuidString := uuidData.String()
	CheckErr(err)
	return uuidString, err
}
