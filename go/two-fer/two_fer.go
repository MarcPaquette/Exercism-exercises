//twofer provides a function that modifies names into a sentence

package twofer

// ShareWith returns string formated with provided name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
