package generate

import (
	"errors"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/root"
	"github.com/calebcase/gomnia/lib/dist/scaled"
	"github.com/calebcase/gomnia/lib/dist/truncated"
)

var (
	Cmd = &cobra.Command{
		Use:   "generate",
		Short: "commands for data generation",
		Long:  "Commands for data generation.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			signal.Ignore(syscall.SIGPIPE)

			return
		},
	}

	Bookend      bool
	ScaledMin    float64
	ScaledMax    float64 = 1
	TruncatedMin float64
	TruncatedMax float64 = 1
)

func init() {
	root.Cmd.AddCommand(Cmd)

	flags := Cmd.PersistentFlags()
	flags.BoolVar(&Bookend, "bookend", Bookend, "force the inclusion of the min and max values in the output")
	flags.Float64Var(&ScaledMin, "min", ScaledMin, "scale output to this minimum value")
	flags.Float64Var(&ScaledMax, "max", ScaledMax, "scale output to this maximum value")
	flags.Float64Var(&TruncatedMin, "truncate-min", TruncatedMin, "truncate output to this minimum value")
	flags.Float64Var(&TruncatedMax, "truncate-max", TruncatedMax, "truncate output to this maximum value")
}

func Sample(dist distuv.Rander) (err error) {
	defer func() {
		if errors.Is(err, syscall.EPIPE) {
			err = nil
		}
	}()

	var scaleP, truncateP bool

	rf := Cmd.PersistentFlags()
	switch {
	case rf.Changed("min"):
		scaleP = true
	case rf.Changed("max"):
		scaleP = true
	case rf.Changed("truncate-min"):
		truncateP = true
	case rf.Changed("truncate-max"):
		truncateP = true
	}

	var BookendMin, BookendMax float64

	if truncateP {
		if TruncatedMin >= TruncatedMax {
			return errors.New("truncated min >= truncated max, but must be less than max.")
		}

		truncatableDist, ok := dist.(truncated.TruncatableDistribution)
		if !ok {
			return errors.New("Distribution is not truncatable.")
		}

		dist = truncated.New(truncatableDist, TruncatedMin, TruncatedMax, nil)

		BookendMin = TruncatedMin
		BookendMax = TruncatedMax
	}

	if scaleP {
		if ScaledMin >= ScaledMax {
			return errors.New("scaled min >= truncated max, but must be less than max.")
		}

		dist = scaled.New(dist, ScaledMin, ScaledMax)

		BookendMin = ScaledMin
		BookendMax = ScaledMax
	}

	if Bookend {
		_, err = fmt.Println(BookendMin)
		if err != nil {
			return
		}

		_, err = fmt.Println(BookendMax)
		if err != nil {
			return
		}
	}

	for sample := dist.Rand(); ; sample = dist.Rand() {
		_, err = fmt.Println(sample)
		if err != nil {
			return
		}
	}

	return
}
