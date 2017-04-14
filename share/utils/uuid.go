package utils

import (
	"math/rand"

	"github.com/satori/go.uuid"
)

// GenerateUUID create uuid
func GenerateUUID() uuid.UUID {
	return uuid.NewV4()
}

// GenerateUUIDString create uuid string
func GenerateUUIDString() string {
	return GenerateUUID().String()
}

// GetUUIDFromString convert uuid string to uuid.UUID
func GetUUIDFromString(uuidStr string) (uuid.UUID, error) {
	return uuid.FromString(uuidStr)
}

// ConvertUUIDToString convert uuid.UUID to uuid string
func ConvertUUIDToString(ouuid uuid.UUID) string {
	return ouuid.String()
}

// GenerateAccessToken generate access token
func GenerateAccessToken(param ...interface{}) string {
	return GenerateUUIDString()
}

// RandomStr random count length string
func RandomStr(count int) string {
	if count <= 0 {
		return ""
	}

	charTemps := "0123456789abcdefghigklmnopqrstuvwxyz_"
	strBytes := make([]byte, count)
	for i := 0; i < count; i++ {
		char := charTemps[rand.Intn(len(charTemps))]
		strBytes[i] = char
	}

	return string(strBytes)
}
