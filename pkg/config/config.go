package config

import (
	"github.com/joho/godotenv"
	// "log"
	"os"
)

var Version string = "v1.0"
var CREATOR_NAME string = "Adhitya Rachman H"
var CREATOR, PRODUCT_ID, PRODUCT = GetEnvVal()
var MODE = os.Getenv("MODE")
var allowOrigin []string = []string{"http://127.0.0.1:5500", "http://localhost:4200"}
var allowMethods []string = []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"}
var allowheaders []string = []string{"Content-Type", "Authorization"}

// 	"X-Arh",
// 	"Accept",
// 	"Authorization",
// 	"Content-Type",
// 	"Content-Length",
// 	"X-CSRF-Token",
// 	"Token",
// 	"session",
// 	"Origin",
// 	"Host",
// 	"Connection",
// 	"Accept-Encoding",
// 	"Accept-Language",
// 	"X-Requested-With",
// }
var Debug bool = true

func GetEnvVal() (string, string, string) {
	godotenv.Load()
	c := os.Getenv("CREATOR")
	pi := os.Getenv("PRODUCT_ID")
	p := os.Getenv("PRODUCT")
	return c, pi, p
}

func GetCorsConfig() ([]string, []string, []string, bool) {
	if MODE == "PROD" {
		allowOrigin = []string{"https://cv-arh.web.app"}
		allowMethods = []string{"POST", "GET"}
		Debug = false
	}
	return allowOrigin, allowMethods, allowheaders, Debug
}
