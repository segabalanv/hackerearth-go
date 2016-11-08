package main

import "fmt"

func main() {
	var n int
	var str string 
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &str)
	symbolCount := make(map[rune]int)
	for i := 0; i < n; i++ {
		symbolCount[rune(str[i])]++
	}
	min := 1000000
	for k, v := range symbolCount {
		switch k {
		case 'h', 'a', 'e', 'r':
			if min > v/2 {
				min = v / 2
			}
		case 'c', 'k', 't':
			if min > v {
				min = v
			}
		}
	}
	fmt.Println(min)
}