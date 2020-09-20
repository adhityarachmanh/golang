package utils

import (

	// "fmt"
	// "bufio"

	"arh/pkg/config"
	"arh/pkg/models"

	"encoding/json"

	"strings"
)

func GenNewKey() (map[int]int, []int) {

	m1 := ""
	for i := 32; i < 127; i++ {
		m1 += string(rune(i))
	}
	t := m1
	m2 := ""
	l := len(m1)

	for i := 0; i < l; i++ {
		c := randomChoice(t)
		m2 += c
		t = strings.ReplaceAll(t, c, "")
	}
	d1 := map[int]int{}
	d2 := map[int]int{}
	o1 := 0
	o2 := 0
	var a []int
	for i := 0; i < l; i++ {
		o1 = int(m1[i])
		o2 = int(m2[i])
		a = append(a, o2)
		d1[o1] = o2
		d2[o2] = o1
	}
	a = _encKey(a)
	// fmt.Println("new key :", d1)
	// fmt.Println("new key :", a)
	return d1, a
}

func _encKey(arr []int) []int {

	s := config.CREATOR + config.PRODUCT_ID + config.PRODUCT
	lenS := len(s)
	var k []int
	l := len(arr)
	for i := 0; i < l; i++ {
		k = append(k, int(s[i%lenS])^arr[i])
	}
	return k
}

func Enc(s string, key map[int]int) string {

	return Translate(s, key)
}
func getDicD(arr []int) map[int]int {
	s := config.CREATOR + config.PRODUCT_ID + config.PRODUCT
	lenS := len(s)
	k := 0
	key := make(map[int]int)
	l := len(arr)
	for i := 0; i < l; i++ {
		k = int(s[i%lenS]) ^ arr[i]
		key[k] = i + 32
	}
	return key
}
func Dec(s string, arrKey []int) string {
	key := getDicD(arrKey)
	return Translate(s, key)
}

func Translate(sn string, d map[int]int) string {
	out := []rune(sn)
	// d = map[int]int{97: 96}
	var c string
	for i := 0; i < len(out); i++ {
		aw := out[i]
		// fmt.Println("awal", aw)
		out[i] = rune(d[int(aw)])
		// fmt.Println("change", d[int(aw)], "->", string(out[i]))
		c += string(out[i])
	}

	// for key, value := range d {
	// 	out[key] = rune(value)
	// 	fmt.Println(out)
	// }

	return c

}

func GetDecData(filename string, password string, d interface{}) error {
	var err error
	if password == "" {
		password = HashAndSalt(config.CREATOR + config.PRODUCT_ID + config.PRODUCT)
	}
	var Key models.KeySchema
	var conn string
	JsonLoads(string(DecryptFile("key/key", "")), &Key)
	dataFile := DecryptFile("key/"+filename, password)
	err = json.Unmarshal(dataFile, &conn)
	if err != nil {
		return err
	}
	data := Dec(conn, Key.Dec)
	json.Unmarshal([]byte(data), &d)
	if err != nil {
		return err
	}
	return err
}
