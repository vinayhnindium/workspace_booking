package config

import (
	"fmt"
	"os"
)

func GetServerPort() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
}
