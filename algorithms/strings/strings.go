package strings

// GetAllSubstrings returns all sequent sub-strings from string
func GetAllSubstrings(str string) []string {
	res := make([]string, 0)
	// starting point
	for i := 0; i < len(str); i++ {
		// length of a substring
		for j := 1; j <= len(str)-i; j++ {
			res = append(res, str[i:i+j])
		}
	}
	return res
}
