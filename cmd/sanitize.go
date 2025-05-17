package cmd

import (
	"github.com/spf13/cobra"
	"thoth/internal/hashes"
	"thoth/internal/util"
)

// Variables for flags
var mode int
var ifilePath string
var ofilePath string

var saniCmd = &cobra.Command{
	Use:   "sanitize",
	Short: "Sanitize Hash File",
	Long:  `Sanitizes hash file based on hash module to be used.`,
	Run: func(cmd *cobra.Command, args []string) {

		if ifilePath == "" {
			util.Red("You must provide a file to sanitize.\n")
			cmd.Usage()
			return
		}

		hashes.HashCheck(mode, "asdf")

		// If ofilePath is "" then we will just print out to console
	},
}

func init() {
	rootCmd.AddCommand(saniCmd)

	// Define custom multi-character short flags
	saniCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Specify Hash Algorithm Utilize: https://hashcat.net/wiki/doku.php?id=example_hashes")
	saniCmd.Flags().StringVarP(&ifilePath, "file", "f", "", "Path to the input file")
	saniCmd.Flags().StringVarP(&ofilePath, "output", "o", "", "Path to the output file")
}
