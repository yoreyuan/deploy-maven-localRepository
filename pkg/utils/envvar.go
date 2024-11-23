package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetEnvOrDefault(envVar, defaultValue string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	}
	return defaultValue
}

func GetIntValueFromEnv(envVar string, defaultValue int) int {
	if v, ok := os.LookupEnv(envVar); ok {
		if v != "" {
			if result, err := strconv.Atoi(v); err == nil {
				return result
			}
		}
	}
	return defaultValue
}

// ReadBool returns the boolean value of an envvar of the given name.
func ReadBool(envVarName string) bool {
	envVar := GetEnvOrDefault(envVarName, "false")
	return strings.TrimSpace(strings.ToLower(envVar)) == "true"
}
