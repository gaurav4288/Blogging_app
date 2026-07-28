package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	SessionTimeout    time.Duration
	MaxFailedAttempts int
	LockoutDuration   time.Duration
}

func Load() *Config {
	return &Config{
		DBHost:            getEnv("DB_HOST", "postgres"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBUser:            getEnv("DB_USER", "cliuser"),
		DBPassword:        getEnv("DB_PASSWORD", "clipassword"),
		DBName:            getEnv("DB_NAME", "clilogin"),
		SessionTimeout:    getDurationMinutes("SESSION_TIMEOUT", 30),
		MaxFailedAttempts: getEnvInt("MAX_FAILED_ATTEMPTS", 5),
		LockoutDuration:   getDurationMinutes("LOCKOUT_DURATION", 15),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}

func getDurationMinutes(key string, defaultMin int) time.Duration {
	mins := getEnvInt(key, defaultMin)
	return time.Duration(mins) * time.Minute
}
