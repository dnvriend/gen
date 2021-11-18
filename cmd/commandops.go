package cmd

import "github.com/spf13/cobra"

type CommandOps struct {
	Cmd *cobra.Command
}

func (c CommandOps) GetBoolParam(name string) bool {
	return c.GetBoolParamDefault(name, false)
}

func (c CommandOps) GetBoolParamDefault(name string, defaultValue bool) bool {
	v, err := c.Cmd.Flags().GetBool(name)
	if err != nil {
		return defaultValue
	}
	return v
}

func (c CommandOps) GetStringParam(name string) string {
	v, err := c.Cmd.Flags().GetString(name)
	cobra.CheckErr(err)
	return v
}

func (c CommandOps) GetStringArrayParam(name string) []string {
	v, err := c.Cmd.Flags().GetStringArray(name)
	cobra.CheckErr(err)
	return v
}

func (c CommandOps) GetIntParam(name string) int {
	v, err := c.Cmd.Flags().GetInt(name)
	cobra.CheckErr(err)
	return v
}
