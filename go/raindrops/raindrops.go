/*Package raindrops converts a number to a string, the contents of which depend on the number's factors.

- If the number has 3 as a factor, output 'Pling'.
- If the number has 5 as a factor, output 'Plang'.
- If the number has 7 as a factor, output 'Plong'.
- If the number does not have 3, 5, or 7 as a factor,
  just pass the number's digits straight through.
*/
package raindrops

import "strconv"

type factorSet struct {
	factor    int
	dropsound string
}

func (f factorSet) appender(number int, output string) string {
	if number%f.factor == 0 {
		output = output + f.dropsound
	}

	return output
}

//Convert a number into a string
func Convert(number int) string {

	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return strconv.Itoa(number)
	}

	var factorList [3]factorSet
	factorList[0] = factorSet{3, "Pling"}
	factorList[1] = factorSet{5, "Plang"}
	factorList[2] = factorSet{7, "Plong"}

	var output string

	for _, raindrop := range factorList {
		output = raindrop.appender(number, output)
	}

	return output
}
