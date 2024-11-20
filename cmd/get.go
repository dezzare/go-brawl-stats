package cmd

import (
	"fmt"
	"os"

	apiClient "github.com/dezzare/go-brawl-stats/api/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data from Supercell",
	Long:  "This command will call Supercell API in order to return the desired value",
	Run: func(cmd *cobra.Command, args []string) {
		c := apiClient.New()
		tag := "%23V0CJ2J"
		out, err := c.GetPlayer(tag)
		if err != nil {
			fmt.Print(err)
		}
		os.Stdout.Write(out)
	},
}
