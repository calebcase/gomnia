package summarize

import (
	"github.com/spf13/cobra"

	"github.com/calebcase/gomnia/cmd/root"
)

var (
	Cmd = &cobra.Command{
		Use:   "summarize",
		Short: "commands for summarizing data",
		Long:  "Commands for summarizing data.",
	}
)

func init() {
	root.Cmd.AddCommand(Cmd)
}
