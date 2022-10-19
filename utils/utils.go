package utils

import "strconv"

// get last index
func GetLastIndex(splits []string) string {
	return splits[len(splits)-1]
}

// string to int
// type convert
func SplitToInt(s string) int {
	intStr, err := strconv.Atoi(s)
	ErrorCheck(err)
	return intStr
}

// error check
func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
