package truncated

import (
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type TruncatableDistribution interface {
	CDF(x float64) float64
	Quantile(p float64) float64
}

// https://www.jstatsoft.org/article/view/v016c02/v16c02.pdf
type Truncated struct {
	Distribution TruncatableDistribution

	Min float64
	Max float64

	Src rand.Source

	cdfMin *float64
	cdfMax *float64
}

func (t *Truncated) cache() {
	if t.cdfMin == nil {
		cdfMin := t.Distribution.CDF(t.Min)
		t.cdfMin = &cdfMin
	}

	if t.cdfMax == nil {
		cdfMax := t.Distribution.CDF(t.Max)
		t.cdfMax = &cdfMax
	}
}

func (t Truncated) CDF(x float64) float64 {
	t.cache()

	cdfMaxMin := t.Distribution.CDF(math.Max(math.Min(x, t.Max), t.Min))

	cdfMin := *t.cdfMin
	cdfMax := *t.cdfMax

	return (cdfMaxMin - cdfMax) / (cdfMax - cdfMin)
}

func (t Truncated) Quantile(p float64) float64 {
	t.cache()

	cdfMin := *t.cdfMin
	cdfMax := *t.cdfMax

	return t.Distribution.Quantile(cdfMin + p*(cdfMax-cdfMin))
}

func (t Truncated) Rand() float64 {
	t.cache()

	p := distuv.Uniform{
		Min: 0,
		Max: 1,

		Src: t.Src,
	}.Rand()

	return t.Quantile(p)
}
