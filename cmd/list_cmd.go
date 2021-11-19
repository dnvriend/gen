package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dnvriend/gen/list"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "generates List[T] code",
	Run: func(_cmd *cobra.Command, args []string) {
		cmd := ToCobraCommand(_cmd)
		packageName := cmd.GetStringParam("package")
		typeName := cmd.GetStringParam("type")
		mapTo := cmd.GetStringArrayParam("mapto")
		foldMapTo := cmd.GetStringArrayParam("foldmapto")
		imports := cmd.GetStringArrayParam("import")

		generated := list.Generate(packageName, typeName, mapTo, foldMapTo, imports)
		switch {
		case cmd.GetBoolParam("stdout"):
			fmt.Println(generated)
		default:
			fileTypeName := strings.ReplaceAll(typeName, "*", "")
			fileTypeName = strings.ReplaceAll(fileTypeName, ".", "")
			dir, err := filepath.Abs(".")
			cobra.CheckErr(err)
			fileName := fmt.Sprintf("%v/%v_list.go", dir, strings.ToLower(fileTypeName))
			err = SaveFile(fileName, []byte(generated))
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("package", "p", "", "the package name")
	listCmd.Flags().StringP("type", "t", "", "the type name")
	listCmd.Flags().StringArrayP("mapto", "m", []string{}, "generate MapTo[T] methods")
	listCmd.Flags().StringArrayP("foldmapto", "f", []string{}, "generate FoldMapTo[T] methods")
	listCmd.Flags().StringArrayP("import", "i", []string{}, "add extra import")
	listCmd.Flags().BoolP("stdout", "s", false, "print to stdout")
}
