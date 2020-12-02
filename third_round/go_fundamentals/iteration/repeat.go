package iteration

//Repeat function to repeat a character many times requested
func Repeat(char string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += char
	}
	return
}
