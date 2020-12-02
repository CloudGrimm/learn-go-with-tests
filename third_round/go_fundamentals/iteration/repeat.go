package iteration

//Repeat function to repeat a character many times requested
func Repeat(char string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += char
	}
	return
}
