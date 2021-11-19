package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/dnvriend/gen/rangeiter"
	"github.com/spf13/cobra"
)

var rangeCmd = &cobra.Command{
	Use:   "range",
	Short: "generates Range code",
	Run: func(_cmd *cobra.Command, args []string) {
		cmd := ToCobraCommand(_cmd)
		packageName := cmd.GetStringParam("package")
		generated := rangeiter.Generate(packageName)
		switch {
		case cmd.GetBoolParam("stdout"):
			fmt.Println(generated)
		default:
			dir, err := filepath.Abs(".")
			cobra.CheckErr(err)
			fileName := fmt.Sprintf("%v/range_type.go", dir)
			err = SaveFile(fileName, []byte(generated))
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rangeCmd)
	rangeCmd.Flags().StringP("package", "p", "", "the package name")
	rangeCmd.Flags().BoolP("stdout", "s", false, "print to stdout")
}
