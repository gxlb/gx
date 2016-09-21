package strings

import (
	xrand "github.com/vipally/gx/math/rand"
)

var (
	rnd = xrand.NewRand32()
)

//rule of string
type Rule uint

const (
	RULE_LIMIT_LEN Rule = 1 << (iota + 16)
	RULE_NUMBER
	RULE_LETTER
	RULE_LETTER_UPPER
	RULE_LETTER_LOWER
	RULE_SYMBOL

	RULE_NORMAL      = RULE_NUMBER | RULE_LETTER
	RULE_STRONG      = Rule(10) | RULE_NUMBER | RULE_LETTER_UPPER | RULE_LETTER_LOWER | RULE_SYMBOL
	RULE_NONE   Rule = 0

	upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower  = "abcdefghijklmnopqrstuvwxyz"
	number = "1234567890"
	sign   = "_#@!$$&*[]{}<>?/\\~"
)

//Generate a rand string with gived rule and length
func Rand(rule Rule, length int) string {
	b := make([]byte, length, length)
	for i := 0; i < length; i++ {
		b[i] = upper[rnd.RandMax(uint32(len(upper)))]
	}
	return string(b)
}

//get rule of a string
func GetRule(str string) Rule {
	return RULE_NONE
}
