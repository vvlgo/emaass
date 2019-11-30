package utils

import "github.com/satori/go.uuid"

func GetUUID() (uuids string, err error) {
	u1 := uuid.Must(uuid.NewV1(), err)
	return u1.String(), err
}
