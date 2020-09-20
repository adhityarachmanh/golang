package connstrpg

import (
	"arh/pkg/models"
	"arh/pkg/utils"
)

func Decrypt(filename string, filePassword string) {
	utils.ClearCMD()
	utils.PrintFigure("Connstrpg", "")
	if valid := utils.FileExists("key/key"); !valid {
		Generate()
	}
	utils.LoggerService.Warnln("Nama File :", filename)
	filePassword = utils.HashAndSalt(filePassword)
	var Key models.KeySchema
	utils.JsonLoads(string(utils.DecryptFile("key/key", "")), &Key)

	decFile := utils.DecryptFile("key/"+filename, filePassword)
	var newDec string
	utils.JsonLoads(string(decFile), &newDec)
	utils.LoggerService.Successln("Enc :", newDec)
	utils.LoggerService.Successln("Dec :", utils.Dec(newDec, Key.Dec))
}
