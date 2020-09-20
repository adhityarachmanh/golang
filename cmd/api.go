package cmd

import (
	// "arh/pkg/generates"
	"arh/api"
	"arh/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.PersistentFlags().IntVarP(&httpPort, "port", "p", 5000, "HTTP port")
}

var apiCmd = &cobra.Command{
	Example: `api serve`,
	Use:     "api",
	Short:   "Serve API",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")

		if port == "" {
			port = "21500"
		}
		utils.ClearCMD()
		utils.PrintFigure("API", "")

		app := api.AppSchema{}
		app.Initialize()
		app.Run(":" + strconv.Itoa(httpPort))
	},
}
