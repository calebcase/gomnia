package generate

import (
	"fmt"

	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/root"
	"github.com/calebcase/gomnia/lib/dist/scaled"
)

var (
	Cmd = &cobra.Command{
		Use:   "generate",
		Short: "commands for data generation",
		Long:  "Commands for data generation.",
	}

	Bookend bool
	Min     float64
	Max     float64 = 1
)

func init() {
	root.Cmd.AddCommand(Cmd)

	flags := Cmd.PersistentFlags()
	flags.BoolVar(&Bookend, "bookend", Bookend, "force the inclusion of the min and max values in the output")
	flags.Float64Var(&Min, "min", Min, "scale output to this minimum value")
	flags.Float64Var(&Max, "max", Max, "scale output to this maximum value")
}

func Sample(dist distuv.Rander) (err error) {
	scaleP := false

	rf := Cmd.PersistentFlags()
	switch {
	case rf.Changed("min"):
		scaleP = true
	case rf.Changed("max"):
		scaleP = true
	}

	if scaleP {
		if Min >= Max {
			root.Log.Error("min is >= max, but must be less than max.", log15.Ctx{
				"min": Min,
				"max": Max,
			})
		}

		dist = scaled.New(dist, Min, Max)

		if Bookend {
			fmt.Println(Min)
			fmt.Println(Max)
		}
	}

	for sample := dist.Rand(); ; sample = dist.Rand() {
		fmt.Println(sample)
	}

	return
}
