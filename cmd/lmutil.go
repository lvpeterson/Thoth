package cmd

import (
	"fmt"
	"github.com/lvpeterson/Thoth/internal/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var crackedHashLM map[string]string
var sanitizedFilePath string

var lmutilCmd = &cobra.Command{
	Use:   "lmutil",
	Short: "LM Hash Utility",
	Long:  `Takes LM cracked hashes and combines them and looks for them in the original hashes. Additionally creates wordlist to utilize against NTLM`,
	Run: func(cmd *cobra.Command, args []string) {
		crackedHashLM = make(map[string]string)
		// Process Cracked Hashes
		content, err := os.ReadFile(ifilePath)
		if err != nil {
			util.Red("Error reading in cracked file")
			return
		}
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			hash := strings.Split(line, ":")
			crackedHashLM[strings.TrimSpace(hash[0])] = strings.TrimSpace(hash[1])
		}

		// Process Uncracked Sanitized Hashes and Create Wordlist for NTLM
		// Create Wordlist Path
		currentDir, err := os.Getwd()
		fileName := fmt.Sprintf("wordlist.txt")
		wordlistFilePath := filepath.Join(currentDir, fileName)
		wordlistFile, err := os.Create(wordlistFilePath)
		if err != nil {
			util.Red("Error creating wordlist file")
			return
		}
		defer wordlistFile.Close()

		file, err := os.Create(ofilePath)
		if err != nil {
			util.Red("Error creating file")
			return
		}
		defer file.Close()

		content, err = os.ReadFile(sanitizedFilePath)
		if err != nil {
			util.Red("Error reading in sanitized file")
			return
		}
		lines = strings.Split(string(content), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			cracked := getFullString(line)
			combinations := generateCaseCombinations(cracked)
			for _, combo := range combinations {
				if _, err := fmt.Fprintf(wordlistFile, "%s\n", combo); err != nil {
					util.Red("Error writing to wordlist file")
					return
				}
			}

			fullLine := line + " " + cracked
			_, err = fmt.Fprintln(file, fullLine)
			if err != nil {
				util.Red("Error writing to file")
				return
			}

		}
	},
}

func generateCaseCombinations(input string) []string {
	if len(input) == 0 {
		return []string{""}
	}

	subCombinations := generateCaseCombinations(input[1:])
	currentUpper := strings.ToUpper(string(input[0]))
	currentLower := strings.ToLower(string(input[0]))

	result := make([]string, 0, len(subCombinations)*2)
	for _, sub := range subCombinations {
		result = append(result, currentUpper+sub)
		result = append(result, currentLower+sub)
	}

	return result
}

func getFullString(s string) string {
	firstHalf := s[:16]
	secondHalf := s[16:32]

	return crackedHashLM[firstHalf] + crackedHashLM[secondHalf]
}

func init() {
	rootCmd.AddCommand(lmutilCmd)

	// Define custom multi-character short flags
	lmutilCmd.Flags().StringVarP(&ifilePath, "file", "f", "", "Path to the cracked file")
	lmutilCmd.Flags().StringVarP(&sanitizedFilePath, "sanifile", "s", "", "Path to original sanitized file")
	lmutilCmd.Flags().StringVarP(&ofilePath, "output", "o", "", "Path to the output file")

}
