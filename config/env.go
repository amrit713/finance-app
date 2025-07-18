package config

import "os"

func LoadEnv(key string, fallback ...string) string {
	env, ok := os.LookupEnv(key)

	if !ok {
		if len(fallback) > 0 {
			return fallback[0]
		}

		return ""
	}

	return env
}
