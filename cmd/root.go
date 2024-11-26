package cmd

import (
	"github.com/dezzare/go-brawl-stats/configs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-brawl-stats",
	Short: "Brawl Stars Statistics",
	Long:  "CLI application to get Brawl Stars Statistics",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("rootCmd")
	// },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(configs.Load)
}
