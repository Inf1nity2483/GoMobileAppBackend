package config

import "os"

func getStringFromEnv(key string, def string) (res string) {
	res = os.Getenv(key)
	if res == "" {
		return def
	}
	return res
}
