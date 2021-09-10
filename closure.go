package main

import "fmt"

func main() {
	value := 10

	for i := 1; i <= value; i++ {
		if i%2 == 0 {
			fmt.Println(even(i))
		} else {
			fmt.Println(odd(i))
		}
	}

}
func even(i int) string {
	result := func() string {
		even := fmt.Sprintf("even -> %d", i)
		return even
	}
	return result()
}

func odd(i int) string {
	result := func() string {
		even := fmt.Sprintf("odd -> %d", i)
		return even
	}
	return result()
}
