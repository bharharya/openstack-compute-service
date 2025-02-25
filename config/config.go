package config

import (
    "os"
)

type Config struct {
    DBHost      string
    DBUser      string
    DBPassword  string
    DBName      string
    OpenStackAuthURL string
    OpenStackUsername string
    OpenStackPassword string
}

func LoadConfig() Config {
    return Config{
        DBHost: os.Getenv("DB_HOST"),
        DBUser: os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName: os.Getenv("DB_NAME"),
        OpenStackAuthURL: os.Getenv("OS_AUTH_URL"),
        OpenStackUsername: os.Getenv("OS_USERNAME"),
        OpenStackPassword: os.Getenv("OS_PASSWORD"),
    }
}
