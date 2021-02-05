package filter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/calebcase/gomnia/cmd/root"
)

var (
	Min float64
	Max float64

	Cmd = &cobra.Command{
		Use:   "filter",
		Short: "read stdin and copy to stdout input which satisfies the filters",
		Long:  "Read stdin and copy to stdout input which satisfies the filters.",
		RunE: func(command *cobra.Command, args []string) (err error) {
			signal.Ignore(syscall.SIGPIPE)
			defer func() {
				if errors.Is(err, syscall.EPIPE) {
					err = nil
				}
			}()

			var minP, maxP bool

			rf := command.PersistentFlags()

			minP = rf.Changed("min")
			maxP = rf.Changed("max")

			if (minP || maxP) == false {
				return errors.New("No filters specified. At least one filter type must be applied.")
			}

			scanner := bufio.NewScanner(os.Stdin)

			var v float64
			for scanner.Scan() {
				v, err = strconv.ParseFloat(scanner.Text(), 64)
				if err != nil {
					return
				}

				if minP && v < Min {
					continue
				}

				if maxP && v > Max {
					continue
				}

				_, err = fmt.Printf("%g\n", v)
				if err != nil {
					return
				}
			}
			err = scanner.Err()
			if err != nil {
				return
			}

			return
		},
	}
)

func init() {
	root.Cmd.AddCommand(Cmd)

	flags := Cmd.PersistentFlags()

	flags.Float64Var(&Min, "min", Min, "drop values less than min (inclusive)")
	flags.Float64Var(&Max, "max", Max, "drop values more than max (inclusive)")
}
