package strings

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
)

//Generate a rand string with gived rule and length
func Rand(rule Rule, length uint) string {
	return ""
}

//get rule of a string
func GetRule(str string) Rule {
	return RULE_NONE
}
