package main

func vStrJoin(sep string, strs ...string) string {
	var finalStr string
	for i, str := range strs {
		// Add the separator only between strings, not after the last one
		if i > 0 {
			finalStr += sep
		}
		finalStr += str
	}
	return finalStr
}
