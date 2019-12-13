package limit

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/calebcase/gomnia/cmd/root"
)

var (
	Min float64
	Max float64

	Sum   float64
	Count uint64

	Cmd = &cobra.Command{
		Use:   "limit",
		Short: "read stdin and copy to stdout until one of the limits are reached",
		Long:  "Read stdin and copy to stdout until one of the limits are reached.",
		RunE: func(command *cobra.Command, args []string) (err error) {
			var minP, maxP, sumP, countP bool

			rf := command.PersistentFlags()

			minP = rf.Changed("min")
			maxP = rf.Changed("max")
			sumP = rf.Changed("sum")
			countP = rf.Changed("count")

			if (minP || maxP || sumP || countP) == false {
				return errors.New("No limits specified. At least one limit type must be applied.")
			}

			scanner := bufio.NewScanner(os.Stdin)

			var sum float64
			var count uint64

			var v float64
			for scanner.Scan() {
				v, err = strconv.ParseFloat(scanner.Text(), 64)
				if err != nil {
					return
				}

				if minP && v < Min {
					break
				}

				if maxP && v > Max {
					break
				}

				if sumP {
					sum += v

					if sum > Sum {
						break
					}
				}

				fmt.Printf("%g\n", v)

				if countP {
					count += 1

					if count >= Count {
						break
					}
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

	flags.Float64Var(&Min, "min", Min, "stop before a value is less than min (inclusive)")
	flags.Float64Var(&Max, "max", Max, "stop before a value is more than max (inclusive)")

	flags.Float64Var(&Sum, "sum", Sum, "stop before the summation of values is greater than sum (inclusive)")
	flags.Uint64Var(&Count, "count", Count, "stop before the number of values is greater than count (inclusive)")
}
