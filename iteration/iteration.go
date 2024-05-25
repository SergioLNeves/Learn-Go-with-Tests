package iteration

func Repeat(repeatCounting int, character string) string {
	var repeated string
	for i := 0; i < repeatCounting; i++ {
		repeated += character
	}

	return repeated
}
