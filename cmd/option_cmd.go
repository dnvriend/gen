package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dnvriend/gen/option"
	"github.com/spf13/cobra"
)

var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "generates option[T] code",
	Run: func(_cmd *cobra.Command, args []string) {
		cmd := ToCobraCommand(_cmd)
		packageName := cmd.GetStringParam("package")
		typeName := cmd.GetStringParam("type")
		mapTo := cmd.GetStringArrayParam("mapto")
		foldMapTo := cmd.GetStringArrayParam("foldmapto")
		imports := cmd.GetStringArrayParam("import")

		generated := option.Generate(packageName, typeName, mapTo, foldMapTo, imports)
		switch {
		case cmd.GetBoolParam("stdout"):
			fmt.Println(generated)
		default:
			fileTypeName := strings.ReplaceAll(typeName, "*", "")
			fileTypeName = strings.ReplaceAll(fileTypeName, ".", "")
			dir, err := filepath.Abs(".")
			cobra.CheckErr(err)
			fileName := fmt.Sprintf("%v/%v_option.go", dir, strings.ToLower(fileTypeName))
			err = SaveFile(fileName, []byte(generated))
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(optionCmd)
	optionCmd.Flags().StringP("package", "p", "", "the package name")
	optionCmd.Flags().StringP("type", "t", "", "the type name")
	optionCmd.Flags().StringArrayP("mapto", "m", []string{}, "generate MapTo[T] methods")
	optionCmd.Flags().StringArrayP("foldmapto", "f", []string{}, "generate FoldMapTo[T] methods")
	optionCmd.Flags().StringArrayP("import", "i", []string{}, "add extra import")
	optionCmd.Flags().BoolP("stdout", "s", false, "print to stdout")
}
