package configs

import "time"

type JwtConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}
