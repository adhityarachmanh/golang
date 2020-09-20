package cmd

import (
	"arh/pkg/config"
	"arh/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
	"os"
	// "strings"
)

var httpPort int
var filename string
var filePassword string
var rootCmd = &cobra.Command{
	Use:     "panic_button [COMMANDS] [OPTIONS]",
	Aliases: []string{"c"},
	// Long:    "arh adalah aplikasi untuk membuat enkripsi dan dekripsi.",
}

func Execute() {
	utils.ClearCMD()
	utils.PrintFigure(config.PRODUCT, "")

	utils.LoggerService.Infoln("CREATOR : ", config.CREATOR)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
