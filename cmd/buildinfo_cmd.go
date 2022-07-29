package cmd

import (
	"fmt"
	"github.com/dnvriend/gen/buildinfo"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var buildinfoCmd = &cobra.Command{
	Use:   "buildinfo",
	Short: "generates build info code",
	Run: func(_cmd *cobra.Command, args []string) {
		checkGitDirExists()
		cmd := ToCobraCommand(_cmd)
		packageName := cmd.GetStringParam("package")
		generated := buildinfo.Generate(ShortCommitHash(), LongCommitHash(), CurrentDateTime(), packageName)
		WriteToFile(generated)
	},
}

func init() {
	rootCmd.AddCommand(buildinfoCmd)
	buildinfoCmd.Flags().StringP("package", "p", "", "the package name")
	buildinfoCmd.Flags().BoolP("stdout", "s", false, "print to stdout")
}

func checkGitDirExists() {
	dir, _ := filepath.Abs("./.git")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("[BuildInfo]: No '.git' directory found at '%v', maybe you need to do a 'git init'?\n", dir)
		os.Exit(1)
	}
}

func ShortCommitHash() string {
	bytes, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	cobra.CheckErr(err)
	return string(bytes)
}

func LongCommitHash() string {
	bytes, err := exec.Command("git", "rev-parse", "HEAD").Output()
	cobra.CheckErr(err)
	return string(bytes)
}

// https://pkg.go.dev/time#pkg-constants
func CurrentDateTime() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func WriteToFile(generated string) {
	dir, err := filepath.Abs(".")
	cobra.CheckErr(err)
	fileName := fmt.Sprintf("%v/buildinfo.go", dir)
	err = SaveFile(fileName, []byte(generated))
	cobra.CheckErr(err)
}
