/*
Package bob uses the function Hey to do the following:
    Bob answers 'Sure.' if you ask him a question. (Is a valid Question ending with ?)
    He answers 'Whoa, chill out!' if you yell at him. (ALL CAPS)
    He answers 'Calm down, I know what I'm doing!' if you yell a question at him. (ALL CAPS and is a Question)
    He says 'Fine. Be that way!' if you address him without actually saying anything. (Provide an empty or whitespace reply)
    He answers 'Whatever.' to anything else (all other scenarios)

*/
package bob

import "strings"

// Hey takes a remark and provides the appropriate response
func Hey(remark string) string {

	cleanedRemark := strings.TrimSpace(remark)
	upperCleanedRemark := strings.ToUpper(cleanedRemark)
	lowerCleanedRemark := strings.ToLower(cleanedRemark)

	if len(cleanedRemark) == 0 {
		return "Fine. Be that way!"
	}

	remarkEnding := cleanedRemark[len(cleanedRemark)-1:]

	switch {
	case remarkEnding == "?" && cleanedRemark == upperCleanedRemark && cleanedRemark != lowerCleanedRemark:
		return "Calm down, I know what I'm doing!"
	case remarkEnding == "?":
		return "Sure."
	case cleanedRemark == upperCleanedRemark && cleanedRemark != lowerCleanedRemark:
		return "Whoa, chill out!"
	case len(cleanedRemark) == 0:
		return "Fine. Be that way!"
	}

	return "Whatever."
}
