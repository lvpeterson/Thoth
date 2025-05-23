package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const Banner = `
████████╗██╗  ██╗ ██████╗ ████████╗██╗  ██╗
╚══██╔══╝██║  ██║██╔═══██╗╚══██╔══╝██║  ██║
   ██║   ███████║██║   ██║   ██║   ███████║
   ██║   ██╔══██║██║   ██║   ██║   ██╔══██║
   ██║   ██║  ██║╚██████╔╝   ██║   ██║  ██║
   ╚═╝   ╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝  ╚═╝
        Keeper of Mysteries
`

// Variables for flags
var mode int
var ifilePath string
var ofilePath string

const (
	ModeNTLM          = 1000
	ModeLM            = 3000
	ModeNetNTLMv1     = 5500
	ModeNetNTLMv2     = 5600
	ModeKerberosTGS   = 13100
	ModeKerberosASREP = 18200
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Thoth",
	Short: "A tool for hash parsing",
	Long:  `Thoth is a Sanitization and Desanitization utility for hashes to be used outside without worry of data leakage.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println(Banner)
	// Disable the default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
