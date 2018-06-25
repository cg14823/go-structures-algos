package algos

// This package contains common string algorithm or problems

// BalancedBracket takes a string composed of brackets and returns true
// if balanced otherwise false. if string contains any other character
// apart from a [](){} it will return false
func BalancedBracket(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	stack := []string{string(str[0])}
	brackets := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	ix := 1
	for ix < len(str) {
		if v, ok := brackets[string(str[ix])]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(str[ix]))
			}
		} else {
			stack = append(stack, string(str[ix]))
		}
		ix++
	}

	return len(stack) == 0
}
