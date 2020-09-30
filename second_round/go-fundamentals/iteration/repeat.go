package iteration

//Repeat function to repeat some set of characters
func Repeat(input string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += input
	}
	return repeated
}
