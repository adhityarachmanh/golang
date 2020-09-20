package config

import (
	"github.com/joho/godotenv"
	// "log"
	"os"
)

func GetEnvVal() (string, string, string) {
	godotenv.Load()
	c := os.Getenv("CREATOR")
	pi := os.Getenv("PRODUCT_ID")
	p := os.Getenv("PRODUCT")
	return c, pi, p
}

var Version string = "v1.0"
var CREATOR, PRODUCT_ID, PRODUCT = GetEnvVal()
var MODE = os.Getenv("MODE")
