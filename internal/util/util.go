package util

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

// Green message for [+]
func Green(message string) {
	color.Green("[+] %s", message)
}

// Red message for [!]
func Red(message string) {
	color.Red("[!] %s", message)
}

// Blue message for [*]
func Blue(message string) {
	color.Blue("[*] %s", message)
}

// Need to pass in a string and a delimitor then return an array of all the components so it can then be parsed afterwards

// SetupProgressBar creates and starts a new progress bar with the specified total value and a width of 50.
func SetupProgressBar(total int) *pb.ProgressBar {
	bar := pb.New(total)
	bar.SetWidth(50)
	bar.Start()
	return bar
}

// RunCommand executes the command specified by name with the given arguments.
func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
