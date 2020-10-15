package utils

import (
	// "arh/pkg/models"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func EdGenkey(tx int) string {
	var t, p, q, l, x, y int
	var k, s, c string
	var a []int

	// fmt.Println("t:", t)
	t = tx
	p = 2
	q = 1
	a = []int{65, 91, 97, 123, 48, 58}
	if t == 0 {
		s = "+/="
	} else if t == 1 {
		s = "-_."
	} else if t == 2 {
		a = []int{97, 123, 63, 91, 48, 59}
	} else if t == 5 {
		s = "!#$%&"
		a = []int{40, 92, 93, 127}
	} else if t == 6 {
		s = "-_."
		a = []int{48, 58, 97, 123}
	} else if t == 7 {
		a = []int{50, 58, 65, 73, 74, 79, 80, 91}
	} else if t == 8 {
		a = []int{48, 58, 65, 91}
	} else if t == 9 {
		c = EdGenkey(10)
		for i := 0; i < 64; i++ {
			s = randomChoice(c)
			c = strings.ReplaceAll(c, s, "")
			k += s
		}
		return k
	} else if t == 10 {
		a = []int{32, 127}
	} else {
		k = "!"
		a = []int{35, 39, 40, 47, 48, 92, 93, 96, 97, 127, 192, 231}
	}

	l = len(a)
	for _, i := range makeRange(0, l, p) {
		// fmt.Println("test:", i)
		x = a[i]
		y = a[i+1]
		if p == 3 {
			q = a[i+2]
		}
		for _, j := range makeRange(x, y, q) {
			k += string(rune(j))
		}
	}
	// fmt.Println("test:", q, x, y)
	return k + s
}

func makeRange(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}
	s := make([]int, 0, 1+(end-start)/step)
	for start <= (end - step) {
		s = append(s, start)
		start += step
	}
	return s
}

func Rank(t int, p int) int {
	for i := 1; i < p; i++ {
		t += t
	}
	return t
}

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		LoggerService.Errorln(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		LoggerService.Errorln(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		LoggerService.Errorln(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		LoggerService.Errorln(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		LoggerService.Errorln(err.Error())
	}
	return plaintext
}

func EncryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(Encrypt(data, passphrase))
}

func DecryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return Decrypt(data, passphrase)
}
