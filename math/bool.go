package math

//bool
type Bool bool

func (me Bool) I() (i int) {
	if bool(me) {
		i = 1
	}
	return
}

func BToI(b Bool) int {
	return Bool(b).I()
}
func ItoB(i int) bool {
	return i != 0
}
