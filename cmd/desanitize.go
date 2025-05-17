package cmd

import "github.com/spf13/cobra"

var deSaniCmd = &cobra.Command{
	Use:   "desanitize",
	Short: "Desanitize Crack File",
	Long:  `Takes cracked results and compares against original hash and combines.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(deSaniCmd)

	// Define custom multi-character short flags
	deSaniCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Specify Hash Algorithm Utilize: https://hashcat.net/wiki/doku.php?id=example_hashes")
	deSaniCmd.Flags().StringVarP(&ifilePath, "file", "f", "", "Path to the input file")
	deSaniCmd.Flags().StringVarP(&ofilePath, "output", "o", "", "Path to the output file")
}
