package utils

import (
	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	uid := uuid.NewV4()
	//return strings.ReplaceAll(uid.String(), "-", ""), nil
	return uid.String()
}
