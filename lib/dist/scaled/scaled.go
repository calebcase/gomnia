package scaled

import (
	"gonum.org/v1/gonum/stat/distuv"
)

type stream struct {
	Min float64
	Max float64
}

type Scaled struct {
	Distribution distuv.Rander

	Min float64
	Max float64

	x *stream
}

func New(dist distuv.Rander, min, max float64) (s *Scaled) {
	s = &Scaled{
		Distribution: dist,

		Min: min,
		Max: max,

		x: &stream{},
	}

	var x1, x2 float64

	x1 = s.Distribution.Rand()
	for x2 = s.Distribution.Rand(); x1 == x2; x2 = s.Distribution.Rand() {
	}

	if x1 < x2 {
		s.x.Min = x1
		s.x.Max = x2
	} else {
		s.x.Min = x2
		s.x.Max = x1
	}

	return
}

func (s Scaled) Rand() float64 {
	var x float64

	for {
		x = s.Distribution.Rand()

		if x < s.x.Min {
			s.x.Min = x
		} else if x > s.x.Max {
			s.x.Max = x
		} else {
			break
		}
	}

	p := s.Min + ((x-s.x.Min)*(s.Max-s.Min))/(s.x.Max-s.x.Min)

	return p
}
