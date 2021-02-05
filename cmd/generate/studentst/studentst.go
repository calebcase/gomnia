package studentst

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Mu    float64
	Sigma float64 = 1
	Nu    float64 = 1

	Cmd = &cobra.Command{
		Use:   "students-t",
		Short: "generate samples from the Student's t distribution",
		Long: `Generate samples from the Student's t distribution.

https://en.wikipedia.org/wiki/Student%27s_t-distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.StudentsT{
				Mu:    Mu,
				Sigma: Sigma,
				Nu:    Nu,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Mu, "mu", "m", Mu, "expected value")
	flags.Float64VarP(&Sigma, "sigma", "s", Sigma, "standard deviation")
	flags.Float64VarP(&Nu, "nu", "n", Nu, "degrees of freedom")
}
