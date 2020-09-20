package connstrpg

import (
	"arh/pkg/models"
	"arh/pkg/utils"

	"bufio"

	"fmt"
	"os"
	"strings"
)

func Generate() {
	utils.ClearCMD()
	utils.PrintFigure("Connstrpg", "")
	info, _ := os.Stat("key")
	if valid := info != nil; !valid {
		os.MkdirAll("key", os.ModePerm)
		utils.LoggerService.Warnln("Anda belum memiliki key enc dec")
		GenNewKey()
	} else if valid := utils.FileExists("key/key"); !valid {
		utils.LoggerService.Warnln("Anda belum memiliki key enc dec")
		GenNewKey()
	} else {
		var Key models.KeySchema
		utils.JsonLoads(string(utils.DecryptFile("key/key", "")), &Key)
		utils.LoggerService.Infoln("Enc : ", utils.JsonDumps(Key.Enc)[0:100], "...")
		utils.LoggerService.Infoln("Dec : ", utils.JsonDumps(Key.Dec)[0:100], "...")
		var answer string

		reader := bufio.NewReader(os.Stdin)
		for strings.TrimSpace(answer) == "" {

			utils.LoggerService.Warnln("Ingin generate key baru ? (y/n)")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			answer = strings.Replace(text, "\n", "", -1)
			answer = strings.ToLower(answer)

		}
		if valid := answer == "y" || answer == "n"; !valid {
			utils.ClearCMD()
			Generate()
		}
		if valid := answer == "y"; valid {
			os.RemoveAll("key")
			utils.ClearCMD()
			utils.PrintFigure("Connstrpg", "")
			GenNewKey()
		}
	}

}

func GenNewKey() {
	utils.LoggerService.Warnln("Pembuatan key baru")
	var Key models.KeySchema
	enc, dec := utils.GenNewKey()
	Key.Enc = enc
	Key.Dec = dec
	jsonData := utils.JsonDumps(Key)
	utils.EncryptFile("key/key", []byte(jsonData), "")

	utils.JsonLoads(string(utils.DecryptFile("key/key", "")), &Key)
	utils.LoggerService.Infoln("Enc : ", utils.JsonDumps(Key.Enc)[0:100], "...")
	utils.LoggerService.Infoln("Dec : ", utils.JsonDumps(Key.Dec)[0:100], "...")
	utils.LoggerService.Successln("Berhasil membuat key baru")

}
