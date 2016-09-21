package strings

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	//rnd.Srand(10)
	for i := 0; i < 100; i++ {
		fmt.Println(Rand(RULE_STRONG, 8))
	}
}
