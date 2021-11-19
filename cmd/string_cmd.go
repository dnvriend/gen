package cmd

import (
	"fmt"
	"path/filepath"

	str "github.com/dnvriend/gen/string"
	"github.com/spf13/cobra"
)

var stringCmd = &cobra.Command{
	Use:   "string",
	Short: "generates String code",
	Run: func(_cmd *cobra.Command, args []string) {
		cmd := ToCobraCommand(_cmd)
		packageName := cmd.GetStringParam("package")
		generated := str.Generate(packageName)
		switch {
		case cmd.GetBoolParam("stdout"):
			fmt.Println(generated)
		default:
			dir, err := filepath.Abs(".")
			cobra.CheckErr(err)
			fileName := fmt.Sprintf("%v/string_type.go", dir)
			err = SaveFile(fileName, []byte(generated))
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stringCmd)
	stringCmd.Flags().StringP("package", "p", "", "the package name")
	stringCmd.Flags().BoolP("stdout", "s", false, "print to stdout")
}
