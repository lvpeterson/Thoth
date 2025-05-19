package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"thoth/internal/hashes"
	"thoth/internal/util"
)

var desanFilePath string

const (
	ModeNTLM          = 1000
	ModeLM            = 3000
	ModeNetNTLMv1     = 5500
	ModeNetNTLMv2     = 5600
	ModeKerberosTGS   = 13100
	ModeKerberosASREP = 18200
)

type hashProcessor func([]string) (string, error)

var modeProcessors = map[int]hashProcessor{
	ModeNTLM:          processNTLM,
	ModeLM:            processLM,
	ModeNetNTLMv1:     processNetNTLM,
	ModeNetNTLMv2:     processNetNTLM,
	ModeKerberosTGS:   processKerberosTGS,
	ModeKerberosASREP: processKerberosASREP,
}

var deSaniCmd = &cobra.Command{
	Use:   "desanitize",
	Short: "Desanitize Crack File",
	Long:  `Takes cracked results and compares against original hash and combines.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ifilePath == "" {
			util.Red("You must provide a cracked hash file in order to desanitize.\n")
			cmd.Usage()
			return
		}
		content, err := os.ReadFile(ifilePath)
		if err != nil {
			util.Red("Error reading in cracked hash file")
			return
		}
		if desanFilePath == "" {
			util.Red("You must provide a the original file utilized prior to sanitization (full hash file).\n")
			cmd.Usage()
			return
		}
		origcontent, err := os.ReadFile(desanFilePath)
		if err != nil {
			util.Red("Error reading in original hash file")
			return
		}

		if ofilePath == "" {
			ofilePath = fmt.Sprintf("desanitized_%s.txt", ifilePath)
		}
		file, err := os.Create(ofilePath)
		if err != nil {
			util.Red("Error creating file")
			return
		}
		defer file.Close()

		lines := strings.Split(string(content), "\n")
		for lineNum, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			if !(hashes.HashCheck(mode, line)) {
				util.Red(fmt.Sprintf("Hash at line %d has invalid format for mode %d: %s", lineNum+1, mode, line))
				return
			}
			hashArray := hashes.GetHashArray(mode, line)
			if err := processHash(mode, hashArray, file); err != nil {
				util.Red(fmt.Sprintf("Error processing hash at line %d: %v", lineNum+1, err))
				return
			}
			if err != nil {
				util.Red(fmt.Sprintf("Error writing to file at line %d: %v", lineNum, err))
				return
			}
		}
		util.Green(fmt.Sprintf("Sanitized file saved to: %s", ofilePath))

	},
}

func init() {
	rootCmd.AddCommand(deSaniCmd)

	// Define custom multi-character short flags
	deSaniCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Specify Hash Algorithm Utilize: https://hashcat.net/wiki/doku.php?id=example_hashes")
	deSaniCmd.Flags().StringVarP(&ifilePath, "file", "f", "", "Path to the cracked hashes file")
	deSaniCmd.Flags().StringVarP(&desanFilePath, "desan", "d", "", "Path to the original hash file (prior to sanitization)")
	deSaniCmd.Flags().StringVarP(&ofilePath, "output", "o", "", "Path to the output file")
}
