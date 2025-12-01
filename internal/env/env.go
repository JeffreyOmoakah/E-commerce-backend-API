package env

import (
	"os"
	"strconv"
)

// GetString returns the value of the environment variable with the given key.
// If the variable is not set, it returns the fallback value.
func GetString(key, fallback string ) string {
	val, ok := os.LookupEnv(key)
	
	if !ok{
		return fallback
	}
	return val
}
		
// GetInt returns the value of the environment variable with the given key as an integer.
// If the variable is not set or cannot be parsed as an integer, it returns the fallback value.
func GetInt(key string, fallback int ) int {
	val, ok := os.LookupEnv(key)
	
	if !ok{
		return fallback
	}
	
	ValAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	
	return ValAsInt
}