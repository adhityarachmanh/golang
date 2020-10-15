package connstrpg

import (
	"arh/pkg/models"
	"arh/pkg/utils"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Encrypt(input string, filename string, filePassword string) {
	utils.ClearCMD()
	utils.PrintFigure("Connstrpg", "")
	if valid := utils.FileExists("key/key"); !valid {
		Generate()
	}
	if utils.FileExists("key/" + filename) {
		Decrypt(filename, filePassword)
		utils.LoggerService.Warnln("Nama file " + filename + " sudah digunakan.")
		var answer string
		reader := bufio.NewReader(os.Stdin)
		for answer == "" {
			utils.LoggerService.Warnln("Ingin mereplace file " + filename + " (y/n)?")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			answer = strings.Replace(text, "\n", "", -1)
			answer = strings.ToLower(answer)

		}
		if valid := answer == "y" || answer == "n"; !valid {
			utils.ClearCMD()
			Encrypt(input, filename, filePassword)
		}
		if valid := answer == "y"; !valid {
			os.Exit(1)
		}
		utils.ClearCMD()
		utils.PrintFigure("Connstrpg", "")
	}
	var Key models.KeySchema
	utils.JsonLoads(string(utils.DecryptFile("key/key", "")), &Key)

	input = strings.TrimSpace(input)
	filePassword = utils.HashAndSalt(filePassword)

	if input == "" {

		reader := bufio.NewReader(os.Stdin)
		for input == "" {
			utils.LoggerService.Warnln("Masukkan data yang ingin di encrypsi")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			input = strings.Replace(text, "\n", "", -1)
			input = strings.TrimSpace(input)
		}
	}

	encVal := utils.Enc(input, Key.Enc)
	utils.LoggerService.Infoln("File : ", "key/"+filename)
	utils.LoggerService.Infoln("Input : ", input)

	utils.LoggerService.Successln("Enc : ", encVal)
	endD := utils.JsonDumps(utils.Enc(input, Key.Enc))
	utils.EncryptFile("key/"+filename, []byte(endD), filePassword)
	decFile := utils.DecryptFile("key/"+filename, filePassword)
	var jsonLoads string
	utils.JsonLoads(string(decFile), &jsonLoads)
	utils.LoggerService.Successln("Dec :", utils.Dec(jsonLoads, Key.Dec))

}

func EncryptFile(input string, filename string, filePassword string) {
	utils.ClearCMD()
	utils.PrintFigure("Connstrpg", "")
	if valid := utils.FileExists("key/key"); !valid {
		Generate()
	}
	if utils.FileExists("key/" + filename) {
		Decrypt(filename, filePassword)
		utils.LoggerService.Warnln("Nama file " + filename + " sudah digunakan.")
		var answer string
		reader := bufio.NewReader(os.Stdin)
		for answer == "" {
			utils.LoggerService.Warnln("Ingin mereplace file " + filename + " (y/n)?")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			answer = strings.Replace(text, "\n", "", -1)
			answer = strings.ToLower(answer)

		}
		if valid := answer == "y" || answer == "n"; !valid {
			utils.ClearCMD()
			Encrypt(input, filename, filePassword)
		}
		if valid := answer == "y"; !valid {
			os.Exit(1)
		}
		utils.ClearCMD()
		utils.PrintFigure("Connstrpg", "")
	}
	var Key models.KeySchema
	utils.JsonLoads(string(utils.DecryptFile("key/key", "")), &Key)

	input = strings.TrimSpace(input)
	filePassword = utils.HashAndSalt(filePassword)

	if input == "" {

		reader := bufio.NewReader(os.Stdin)
		for input == "" {
			utils.LoggerService.Warnln("Masukkan path file yang ingin diencrypt")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			input = strings.Replace(text, "\n", "", -1)
			input = strings.TrimSpace(input)
		}
	}
	data, _ := ioutil.ReadFile(input)
	newData := strings.TrimSpace(string(data))
	encVal := utils.Enc(string(newData), Key.Enc)
	utils.LoggerService.Infoln("File : ", "key/"+filename)
	utils.LoggerService.Infoln("Input : ", input)

	utils.LoggerService.Successln("Enc : ", encVal)
	endD := utils.JsonDumps(utils.Enc(string(newData), Key.Enc))
	utils.EncryptFile("key/"+filename, []byte(endD), filePassword)
	decFile := utils.DecryptFile("key/"+filename, filePassword)
	var jsonLoads string
	utils.JsonLoads(string(decFile), &jsonLoads)
	utils.LoggerService.Successln("Dec :", utils.Dec(jsonLoads, Key.Dec))

}
