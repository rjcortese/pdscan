package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rjcortese/pdscan/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pdscan [connection-uri]",
	Short: "Scan your data stores for unencrypted personal data (PII)",
	Long:  "Scan your data stores for unencrypted personal data (PII)",
	Run: func(cmd *cobra.Command, args []string) {
		showData, err := cmd.Flags().GetBool("show-data")
		if err != nil {
			log.Fatal(err)
		}

		showAll, err := cmd.Flags().GetBool("show-all")
		if err != nil {
			log.Fatal(err)
		}

		limit, err := cmd.Flags().GetInt("sample-size")
		if err != nil {
			log.Fatal(err)
		}
		if limit < 1 {
			log.Fatal("Limit must be positive")
		}

		processes, err := cmd.Flags().GetInt("processes")
		if err != nil {
			log.Fatal(err)
		}

		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		} else {
			internal.Main(args[0], showData, showAll, limit, processes)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().Bool("show-data", false, "Show data")
	rootCmd.PersistentFlags().Bool("show-all", false, "Show all matches")
	rootCmd.PersistentFlags().Int("sample-size", 10000, "Sample size")
	rootCmd.PersistentFlags().Int("processes", 1, "Processes")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
