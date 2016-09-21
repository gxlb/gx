package strings

import (
	"fmt"
	"os"
	"testing"
)

func TestRand(t *testing.T) {
	//rnd.Srand(10)
	for i := 0; i < 5; i++ {
		fmt.Println(Rand(RULE_STRONG, 8))
	}
	fmt.Println(os.Getwd())
}
