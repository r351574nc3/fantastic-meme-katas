package string-reverse


func reverseString(s string) string {
	retval := [] rune(s)

	for idx, edx := 0, len(retval) - 1; idx < len(retval) / 2; idx, edx = idx + 1, edx  - 1 {
		retval[idx], retval[edx] = retval[edx], retval[idx] // Reverse characters
	} 

	return string(retval)
}
