package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/dnvriend/gen/float"
	"github.com/spf13/cobra"
)

var floatCmd = &cobra.Command{
	Use:   "float",
	Short: "generates float code",
	Run: func(_cmd *cobra.Command, args []string) {
		cmd := ToCobraCommand(_cmd)
		packageName := cmd.GetStringParam("package")
		generated := float.Generate(packageName)
		switch {
		case cmd.GetBoolParam("stdout"):
			fmt.Println(generated)
		default:
			dir, err := filepath.Abs(".")
			cobra.CheckErr(err)
			fileName := fmt.Sprintf("%v/float_type.go", dir)
			err = SaveFile(fileName, []byte(generated))
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(floatCmd)
	floatCmd.Flags().StringP("package", "p", "", "the package name")
	floatCmd.Flags().BoolP("stdout", "s", false, "print to stdout")
}
