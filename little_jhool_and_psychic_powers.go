package main

import "fmt"

func main() {
	var binNum string
	fmt.Scan(&binNum)
	lastNum := 0
	var counter int
	for _, c := range binNum {
		if c == '1' {
			if lastNum == 0 {
				counter = 1
			} else {
				counter++
			}
			lastNum = 1
		} else {
			if lastNum == 0 {
				counter++
			} else {
				counter = 1
			}
			lastNum = 0
		}
		if counter == 6 {
			break
		}
	}
	if counter == 6 {
		fmt.Println("Sorry, sorry!")
	} else {
		fmt.Println("Good luck!")
	}
}
