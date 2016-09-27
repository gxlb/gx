package math

import (
	"math"
)

const (
	Epslon = 0.000001
)

func Equal(a, b float64) bool { return math.Abs(a-b) < Epslon }
