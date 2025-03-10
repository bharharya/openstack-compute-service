package config

import (
	"log"
)

// Secrets holds sensitive credentials
type Secrets struct {
	OpenStackUsername string
	OpenStackPassword string
	OpenStackAuthURL  string
}

// LoadSecrets loads API keys and secrets from environment variables
func LoadSecrets() *Secrets {
	secrets := &Secrets{
		OpenStackUsername: GetEnv("OS_USERNAME", ""),
		OpenStackPassword: GetEnv("OS_PASSWORD", ""),
		OpenStackAuthURL:  GetEnv("OS_AUTH_URL", ""),
	}

	if secrets.OpenStackUsername == "" || secrets.OpenStackPassword == "" || secrets.OpenStackAuthURL == "" {
		log.Fatal("Missing required OpenStack credentials in environment variables")
	}

	return secrets
}
