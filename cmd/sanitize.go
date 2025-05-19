package cmd

import (
	"fmt"
	"github.com/lvpeterson/Thoth/internal/hashes"
	"github.com/lvpeterson/Thoth/internal/util"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type saniHashProcessor func([]string) (string, error)

var saniModeProcessors = map[int]saniHashProcessor{
	ModeNTLM:          saniProcessNTLM,
	ModeLM:            saniProcessLM,
	ModeNetNTLMv1:     saniProcessNetNTLM,
	ModeNetNTLMv2:     saniProcessNetNTLM,
	ModeKerberosTGS:   saniProcessKerberosTGS,
	ModeKerberosASREP: saniProcessKerberosASREP,
}

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
		content, err := os.ReadFile(ifilePath)
		if err != nil {
			util.Red("Error reading in sanitized file")
			return
		}
		if ofilePath == "" {
			ofilePath = fmt.Sprintf("sanitized_%s.txt", ifilePath)
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
			if err := saniProcessHash(mode, hashArray, file); err != nil {
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
	rootCmd.AddCommand(saniCmd)
	// Define custom multi-character short flags
	saniCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Specify Hash Algorithm Utilize: https://hashcat.net/wiki/doku.php?id=example_hashes")
	saniCmd.Flags().StringVarP(&ifilePath, "file", "f", "", "Path to the input file")
	saniCmd.Flags().StringVarP(&ofilePath, "output", "o", "", "Path to the output file")
}

func saniProcessHash(mode int, hashArray []string, file *os.File) error {
	processor, exists := saniModeProcessors[mode]
	if !exists {
		return fmt.Errorf("unsupported mode: %d", mode)
	}

	result, err := processor(hashArray)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(file, result)
	return err
}

func saniProcessNTLM(hashArray []string) (string, error) {
	return hashArray[3], nil
}

func saniProcessLM(hashArray []string) (string, error) {
	return hashArray[2], nil
}

func saniProcessNetNTLM(hashArray []string) (string, error) {
	hashArray[0] = hashes.GenerateRandomString()
	hashArray[2] = hashes.GenerateRandomString()
	return strings.Join(hashArray, ":"), nil
}

func saniProcessKerberosTGS(hashArray []string) (string, error) {
	hashArray[3] = "*" + hashes.GenerateRandomString()
	hashArray[4] = hashes.GenerateRandomString()
	hashArray[5] = hashes.GenerateRandomString() + "/" + hashes.GenerateRandomString() + "*"
	return strings.Join(hashArray, "$"), nil
}

func saniProcessKerberosASREP(hashArray []string) (string, error) {
	asrepUser := strings.Split(hashArray[3], ":")
	asrepUser[0] = hashes.GenerateRandomString() + "@" + hashes.GenerateRandomString() + "." + hashes.GenerateRandomString()
	hashArray[3] = strings.Join(asrepUser, ":")
	return strings.Join(hashArray, "$"), nil
}
