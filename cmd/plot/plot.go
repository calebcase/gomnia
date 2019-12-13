package plot

import (
	"github.com/spf13/cobra"

	"github.com/calebcase/gomnia/cmd/root"
)

var (
	OutputPath string = "-"
)

func init() {
	root.Cmd.AddCommand(Cmd)

	flags := Cmd.PersistentFlags()
	flags.StringVarP(&OutputPath, "output", "o", OutputPath, "output path")
}

var Cmd = &cobra.Command{
	Use:   "plot",
	Short: "commands for plotting data",
	Long:  "Commands for plotting data.",
}
