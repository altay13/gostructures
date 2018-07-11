package strings

// GetAllSubstrings returns all sequent strings from string
func GetAllSubstrings(str string) []string {
	res := make([]string, 0)
	// starting point
	for i := 0; i < len(str); i++ {
		// length of a substring
		for j := 1; j <= len(str)-i; j++ {
			res = append(res, str[i:i+j])

			// << --- just for more low-level interpretation --- >> //
			// tmp := make([]byte, 0)
			// for z := i; z < j+i; z++ {
			// 	tmp = append(tmp, str[z])
			// }
			// res = append(res, string(tmp))
			// << --- just for more low-level interpretation --- >> //
		}
	}
	return res
}
