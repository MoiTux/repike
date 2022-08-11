// Package repike implements Rob Pike's simple C regex matcher in Go.
package repike

// Match reports whether regexp matches anywhere in text.
func Match(regexp, text string) bool {
	if regexp != "" && regexp[0] == '^' {
		return matchHere(regexp[1:], text)
	}
	for {
		if matchHere(regexp, text) {
			return true
		}
		if text == "" {
			return false
		}
		text = text[1:]
	}
}

// matchHere reports whether regexp matches at beginning of text.
func matchHere(regexp, text string) bool {
	switch {
	case regexp == "":
		return true
	case regexp == "$":
		return text == ""
	case len(regexp) >= 2 && regexp[1] == '*':
		return matchStar(regexp[0], regexp[2:], text)
	case text != "" && (regexp[0] == '.' || regexp[0] == text[0]):
		return matchHere(regexp[1:], text[1:])
	}
	return false
}

// matchStar reports whether c*regexp matches at beginning of text.
func matchStar(c byte, regexp, text string) bool {
	for {
		if matchHere(regexp, text) {
			return true
		}
		if text == "" || (text[0] != c && c != '.') {
			return false
		}
		text = text[1:]
	}
}
