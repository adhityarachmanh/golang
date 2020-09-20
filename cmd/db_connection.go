package cmd

import (
	"fmt"
	// "arh/pkg/generates"

	"arh/connstrpg"
	"arh/pkg/config"
	"arh/pkg/models"
	"arh/pkg/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.PersistentFlags().StringVarP(&filePassword, "password", "p", config.CREATOR+config.PRODUCT_ID+config.PRODUCT, "[OPTIONAL] Password untuk file")
	dbCmd.PersistentFlags().StringVarP(&filename, "set", "s", "", "[REQUIRED] Untuk set file koneksi")
}

var dbCmd = &cobra.Command{
	Example: `
		[DB]
		   -PASSWORD DEFAULT-
			  -> db --set=[FILE NAME]
			  
		   -PASSWORD COSTUME-
		      -> db --password=[PASSWORD FILE] --set=[FILE NAME]
	`,
	Use:     "database",
	Aliases: []string{"db"},
	Short:   "Set database connection",
	Run: func(cmd *cobra.Command, args []string) {
		utils.ClearCMD()
		utils.PrintFigure("DB Connection", "")
		if valid := utils.FileExists("key/key"); !valid {
			connstrpg.Generate()
		}
		if valid := utils.FileExists("key/" + filename); !valid {
			utils.LoggerService.Errorln(fmt.Sprintf("File [%s] tidak ditemukan.", filename))
			os.Exit(1)
		}
		var dbconn models.DatabaseSchema
		passHash := utils.HashAndSalt(filePassword)
		err := utils.GetDecData(filename, passHash, &dbconn)
		if err != nil {
			utils.LoggerService.Errorln("Telah terjadi kesalah decrypsi file")
			os.Exit(1)
		}

		dbc, _ := json.Marshal(dbconn)
		encFileData := utils.DecryptFile("key/"+filename, passHash)

		utils.EncryptFile("key/db", encFileData, passHash)
		utils.LoggerService.Successln(fmt.Sprintf("Koneksi berhasil diubah, dengan file koneksi [%s]", filename))
		utils.LoggerService.Infoln("Data :" + string(dbc))
	},
}
