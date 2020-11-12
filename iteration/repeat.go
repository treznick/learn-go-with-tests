package iteration

func Repeat(character string, length int) string {
	var repeated string
	for i := 0; i < length; i++ {
		repeated += character
	}
	return repeated
}
