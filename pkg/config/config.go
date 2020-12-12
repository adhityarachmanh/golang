package config

import (
	"github.com/joho/godotenv"
	// "log"
	"os"
	"strings"
)

var Version string = "v1.0"

var CREATOR, PRODUCT_ID, PRODUCT, CREATOR_NAME, ALLOW_HOST = GetEnvVal()
var MODE = os.Getenv("MODE")
var allowOrigin []string = []string{"http://localhost:4200"}
var allowMethods []string = []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"}
var allowheaders []string = []string{"Content-Type", "Authorization"}

var Debug bool = true

func GetEnvVal() (string, string, string, string, []string) {
	godotenv.Load()
	var ahh []string
	c := os.Getenv("CREATOR")
	pi := os.Getenv("PRODUCT_ID")
	p := os.Getenv("PRODUCT")
	cn := os.Getenv("CREATOR_NAME")
	ah := strings.TrimSpace(os.Getenv("ALLOW_HOST"))
	ahh = strings.Split(ah, ",")
	return c, pi, p, cn, ahh
}

func GetCorsConfig() ([]string, []string, []string, bool) {
	if MODE == "PROD" {
		allowOrigin = ALLOW_HOST
		allowMethods = []string{"POST", "GET"}
		Debug = false
	}
	return allowOrigin, allowMethods, allowheaders, Debug
}
