package cmd

import "github.com/spf13/cobra"

var lmutilCmd = &cobra.Command{
	Use:   "lmutil",
	Short: "LM Hash Utility",
	Long:  `Takes LM cracked hashes and combines them and looks for them in the original hashes. Additionally creates wordlist to utilize against NTLM`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(lmutilCmd)

	// Define custom multi-character short flags
	lmutilCmd.Flags().StringVarP(&ifilePath, "file", "f", "", "Path to the input file")
	lmutilCmd.Flags().StringVarP(&ofilePath, "output", "o", "", "Path to the output file")
}
