//  Package truncted provides a probability distribution truncater.
package truncated

import (
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type TruncatableDistribution interface {
	distuv.Rander
	distuv.Quantiler

	CDF(x float64) float64
}

// Truncated transforms a distribution by truncating the output between minimum
// and maximum values.
//
// Truncated follows the method described in
// https://www.jstatsoft.org/article/view/v016c02/v16c02.pdf
type Truncated struct {
	Distribution TruncatableDistribution

	Min float64
	Max float64

	Src rand.Source

	cdfMin   float64
	cdfMax   float64
	cdfRange float64
}

var _ distuv.Rander = Truncated{}
var _ distuv.Quantiler = Truncated{}

// New constructs a new truncted distribution with a minimum and maximum range.
func New(dist TruncatableDistribution, min, max float64, src rand.Source) (t Truncated) {
	t = Truncated{
		Distribution: dist,

		Min: min,
		Max: max,

		Src: src,

		cdfMin: dist.CDF(min),
		cdfMax: dist.CDF(max),
	}

	t.cdfRange = t.cdfMax - t.cdfMin

	return
}

// CDF computes the value of the cumulative density function at x.
func (t Truncated) CDF(x float64) float64 {
	cdfMaxMin := t.Distribution.CDF(math.Max(math.Min(x, t.Max), t.Min))

	return (cdfMaxMin - t.cdfMax) / t.cdfRange
}

// Quantile returns the inverse of the cumulative probability distribution.
func (t Truncated) Quantile(p float64) float64 {
	return t.Distribution.Quantile(t.cdfMin + p*t.cdfRange)
}

// Rand returns a random sample drawn from the embedded distribution truncated
// to minimum and maximum values.
func (t Truncated) Rand() float64 {
	p := distuv.Uniform{
		Min: 0,
		Max: 1,

		Src: t.Src,
	}.Rand()

	return t.Quantile(p)
}
