package ascii

import "strings"

func ContainsOnly(string) bool {
	input, _ = InputManagement()
	for i := 0; i < len(input); i++ {
		if !strings.ContainsAny(string(input[i]), "\\n") {
			return false
		}
	}
	return true
}
