package cmd

import (
	// "arh/pkg/generates"
	"arh/api"
	"arh/pkg/utils"
	"github.com/spf13/cobra"
	// "os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.PersistentFlags().IntVarP(&httpPort, "port", "p", 5000, "HTTP port")
	apiCmd.PersistentFlags().StringVarP(&httpHost, "host", "z", "", "HTTP host")

}

var apiCmd = &cobra.Command{
	Example: `api serve`,
	Use:     "api",
	Short:   "Serve API",
	Run: func(cmd *cobra.Command, args []string) {

		utils.ClearCMD()
		utils.PrintFigure("API", "")

		app := api.AppSchema{}
		app.Initialize()
		app.Run(httpHost + ":" + strconv.Itoa(httpPort))
	},
}
