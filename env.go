package main

import (
	"fmt"
	"os"
)

func env(env string) string {
	return os.Getenv(fmt.Sprintf("PLUGIN_%s", env))
}
