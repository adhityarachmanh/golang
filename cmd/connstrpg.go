package cmd

import (
	"arh/connstrpg"
	"arh/pkg/config"
	"arh/pkg/utils"

	"github.com/spf13/cobra"
	"os"
)

var encrypt bool
var encryptFile bool

var decrypt bool
var input string

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().BoolVar(&encrypt, "enc", false, "[REQUIRED] Untuk enccrypt wajib digunakan")
	generateCmd.PersistentFlags().BoolVar(&decrypt, "dec", false, "[REQUIRED] Untuk decrypt wajib digunakan")
	generateCmd.PersistentFlags().BoolVar(&encryptFile, "file", false, "[REQUIRED] Untuk enccrypt wajib digunakan")
	generateCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "[OPTIONAL] Data yg ingin di encrypt <String,int>\nNote: tidak bisa untuk data JSON")
	generateCmd.PersistentFlags().StringVarP(&filePassword, "password", "p", config.CREATOR+config.PRODUCT_ID+config.PRODUCT, "[OPTIONAL] Password untuk file")
	generateCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "[REQUIRED] Untuk encrypt dan decrypt wajib digunakan")

	// rootCmd.AddCommand(encFileCmd)
}

var generateCmd = &cobra.Command{
	Example: `
	[ENC]
	   -DENGAN PASSWORD DEFAULT-
	      -> c -f=[NAMA FILE] --enc 

	   -DENGAN PASSWORD COSTUME-
	      -> c -f=[NAMA FILE] -p=[PASSWORD FILE]  --enc 
	  
	   -DENGAN DATA <STRING,INT>-
		  -> c -f=[NAMA FILE] -i=[DATA] --enc 

	   -DENGAN PATH FILE
	      -> c -f=[NAMA FILE] -i=[DATA] --enc --file

	[DEC]
	   -DENGAN PASSWORD DEFAULT-
	      -> c -f=[NAMA FILE] --dec

	   -DENGAN PASSWORD COSTUME-
	      -> c -f=[NAMA FILE] -p=[PASSWORD FILE] --dec
	`,
	Use:     "connstrpg",
	Aliases: []string{"c", "conn"},
	Short:   "Generate key baru",
	Run: func(cmd *cobra.Command, args []string) {

		if valid := encrypt || decrypt || filename != ""; !valid {
			connstrpg.Generate()
		} else if valid := encrypt && filename != "" && !encryptFile; valid {
			connstrpg.Encrypt(input, filename, filePassword)

		} else if valid := encrypt && filename != "" && encryptFile; valid {
			connstrpg.EncryptFile(input, filename, filePassword)

		} else if valid := decrypt && filename != ""; valid {
			connstrpg.Decrypt(filename, filePassword)

		} else {
			utils.LoggerService.Errorln("Argument tidak valid.")
			os.Exit(1)
		}

	},
}
