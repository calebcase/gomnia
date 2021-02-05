package triangular

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	A float64
	B float64 = 1
	C float64 = 0.5

	Cmd = &cobra.Command{
		Use:   "triangular",
		Short: "generate samples from the triangular distribution",
		Long: `Generate samples from the triangular distribution.

https://en.wikipedia.org/wiki/Triangular_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.NewTriangle(A, B, C, nil))
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&A, "a", "a", A, "best-case estimate (minimum)")
	flags.Float64VarP(&B, "b", "b", B, "worst-case estimate (maximum)")
	flags.Float64VarP(&C, "c", "c", C, "most likely estimate (mode)")
}
