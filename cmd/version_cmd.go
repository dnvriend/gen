package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/dnvriend/gen/typ"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the version",
	Run: func(_cmd *cobra.Command, args []string) {
		fmt.Println(MarshalToString(typ.BuildInfo))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func MarshalToString(v interface{}) string {
	bytes, err := json.Marshal(v)
	cobra.CheckErr(err)
	return string(bytes)
}
