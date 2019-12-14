//  Package scaled provides a probability distribution scaler.
package scaled

import (
	"gonum.org/v1/gonum/stat/distuv"
)

// stream tracks the observed minimum and maximum values from the embedded
// distribution.
type stream struct {
	Min float64
	Max float64
}

// Scaled transforms a distribution by scaling the output between minimum and
// maximum values.
type Scaled struct {
	Distribution distuv.Rander

	Min float64
	Max float64

	observed *stream
}

var _ distuv.Rander = Scaled{}

// New constructs a new scaled distribution with a minimum and maximum range.
func New(dist distuv.Rander, min, max float64) (s Scaled) {
	s = Scaled{
		Distribution: dist,

		Min: min,
		Max: max,

		observed: &stream{},
	}

	var x1, x2 float64

	x1 = s.Distribution.Rand()
	for x2 = s.Distribution.Rand(); x1 == x2; x2 = s.Distribution.Rand() {
	}

	if x1 < x2 {
		s.observed.Min = x1
		s.observed.Max = x2
	} else {
		s.observed.Min = x2
		s.observed.Max = x1
	}

	return
}

// Rand returns a random draw from the embedded distribution rescaled to the
// min and max range.
func (s Scaled) Rand() float64 {
	var x float64

	for {
		x = s.Distribution.Rand()

		if x < s.observed.Min {
			s.observed.Min = x
		} else if x > s.observed.Max {
			s.observed.Max = x
		} else {
			break
		}
	}

	p := s.Min + ((x-s.observed.Min)*(s.Max-s.Min))/(s.observed.Max-s.observed.Min)

	return p
}
