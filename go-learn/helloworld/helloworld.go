package helloworld

import "fmt"

func HelloWorld() (result int) {
	ss, err := fmt.Println("a")
	if err != nil {
		fmt.Println(ss)
	}
	fmt.Println(ss)
	return ss
}

func Sum(set []int) int {
	var result int
	for _, num := range set {
		result += num
	}
	return result
}
