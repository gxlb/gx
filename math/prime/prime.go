package prime

func IsPrime(num uint32) bool {
	if num < 2 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	mod := num % 6
	if mod != 1 && mod != 5 {
		return false
	}
	max := uint32(0) //uint64(math.Floor(math.Sqrt(float64(num))))
	for i := 2; i < len(g_primes); i++ {
		prime := g_primes[i]
		if prime > max {
			break
		}
		if num%prime == 0 {
			return false
		}
	}
	return true
}

func step(shift uint32, val, sqrtVal *uint32) {
	magic := uint32(0x40000000) >> shift
	if (magic + *sqrtVal) <= *val {
		*val -= magic + *sqrtVal
		*sqrtVal = (*sqrtVal >> 1) | magic
	} else {
		*sqrtVal = *sqrtVal >> 1
	}
}

//a fast way to calculate sqrt
func SqrtUint32(val uint32) uint32 {
	var sqrtVal uint32 = 0
	for i := uint32(0); i <= 30; i += 2 {
		step(i, &val, &sqrtVal)
	}
	if sqrtVal < val {
		sqrtVal++
	}
	//sqrtVal <<= (Q_FACTOR) / 2

	return sqrtVal
}
