package cplx

type Cplx64 complex64

func (this *Cplx64) Real() float32 {
	return real(*this)
}
func (this *Cplx64) Imag() float32 {
	return imag(*this)
}

type Cplx128 complex128

func (this *Cplx128) Real() float64 {
	return real(*this)
}
func (this *Cplx128) Imag() float64 {
	return imag(*this)
}
