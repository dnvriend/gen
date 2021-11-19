package cmd

import "github.com/spf13/cobra"

type CobraCommand cobra.Command

func (rcv CobraCommand) getCommand() *cobra.Command {
	return (*cobra.Command)(&rcv)
}

func ToCobraCommand(cmd *cobra.Command) CobraCommand {
	return (CobraCommand)(*cmd)
}

func (rcv CobraCommand) GetBoolParam(name string) bool {
	return rcv.GetBoolParamDefault(name, false)
}

func (rcv CobraCommand) GetBoolParamDefault(name string, defaultValue bool) bool {
	v, err := rcv.getCommand().Flags().GetBool(name)
	if err != nil {
		return defaultValue
	}
	return v
}

func (rcv CobraCommand) GetStringParam(name string) string {
	v, err := rcv.getCommand().Flags().GetString(name)
	cobra.CheckErr(err)
	return v
}

func (rcv CobraCommand) GetStringArrayParam(name string) []string {
	v, err := rcv.getCommand().Flags().GetStringArray(name)
	cobra.CheckErr(err)
	return v
}

func (rcv CobraCommand) GetIntParam(name string) int {
	v, err := rcv.getCommand().Flags().GetInt(name)
	cobra.CheckErr(err)
	return v
}
