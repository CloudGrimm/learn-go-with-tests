package iteration

import "fmt"

func Repeat(input string, repeatedTimes int)  string{
	repeat := ""
	for i := 0; i < repeatedTimes; i++{
		repeat += input
	}
	return repeat
}

func main()  {
	fmt.Println(Repeat("a", 5))
}