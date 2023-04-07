package Iteration

var repeastCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeastCount; i++ {
		repeated += character
	}
	return repeated
}
