package config

import "os"

func JwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
