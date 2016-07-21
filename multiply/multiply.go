package multiply

import (
	"fmt"
	// "bytes"
)

func strtoi(num string) int {
	ZERO := 48
	runes := [] rune(num)
	retval := 0
	
	exp := func(i int) int {
		retval := 1
		if i < 1 {
			return retval
		}
		
		for idx := 1; idx <= i; idx = idx + 1 {
			retval = retval * 10
		} 

		return retval;
	}

	for idx := 0; idx < len(runes); idx = idx + 1 {
		var val int = (int(runes[idx]) - ZERO) * exp(len(runes) - idx - 1)
		retval = retval + val
		fmt.Println("Retval is now ", retval)
	}

	return val
}

func multiply(num1 string, num2 string) string {
	return string(strtoi(num1) * strtoi(num2))
}
