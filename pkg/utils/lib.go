// CREATOR : Adhitya Rachman H
package utils

import (
	"arh/pkg/config"
	"arh/pkg/models"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"sort"
	"time"
)

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

type Block struct {
	Try    func()
	Catch  func(Exception)
	Finish func()
}
type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (cts Block) Go() {
	if cts.Finish != nil {
		defer cts.Finish()
	}
	if cts.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				cts.Catch(r)
			}
		}()
	}
	cts.Try()
}

func RouteAPI(route string) string {
	var URL string
	if config.MODE == "DEV" {
		URL = fmt.Sprintf("/%s/%s/%s", config.Version, "api", route)
	} else {
		URL = fmt.Sprintf("%s/%s/%s", config.Version, "api", route)
		URL = fmt.Sprintf("/%s.arh", Ed.BNE(6, 1).Enc(URL))
	}
	return URL
}

var sR *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func JsonDumps(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func JsonLoads(data string, bind interface{}) {
	json.Unmarshal([]byte(data), &bind)
}

func randomChoice(charset string) string {

	b := make([]byte, 1)
	for i := range b {
		b[i] = charset[sR.Intn(len(charset))]
	}
	return string(b)
}
func ResponseAPIError(c *gin.Context, message string) {
	c.JSON(200, models.ResponseSchema{Message: message, Status: 1, Data: nil, Creator: config.CREATOR_NAME})
}
func ResponseAPI(c *gin.Context, response models.ResponseSchema) {
	response.Creator = config.CREATOR_NAME

	if config.MODE == "PROD" {
		JSONData, _ := json.Marshal(response.Data)
		response.Data = Ed.BNE(6, 2).Enc(string(JSONData))
	}
	c.JSON(200, response)
}

func HashAndSalt(pwd string) string {
	hasher := sha1.New()
	rex := fmt.Sprintf("%s-%s-%s-%s", config.CREATOR, config.PRODUCT, pwd, config.PRODUCT_ID)
	hasher.Write([]byte(rex))
	return hex.EncodeToString(hasher.Sum(nil))
}
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	if hashedPwd != HashAndSalt(plainPwd) {
		return false
	}
	return true
}

func CreateToken(username string) (string, error) {
	var err error
	os.Getenv("CREATOR")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("CREATOR")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ClearCMD() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintFigure(s string, font string) {
	myFigure := figure.NewColorFigure(s, font, "green", true)

	myFigure.Print()
	fmt.Print("\n")
}
func GetMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("")
}
