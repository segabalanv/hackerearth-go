package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	var count [10]int
	str := []rune(input.Text())
	for i := 0; i < len(str); i++ {
		if (str[i]-48 > 0 || str[i]-48 < 10) {
			count[str[i]-48]++
		}
	}
	for i, occ := range count {
		fmt.Println(i, occ)
	}
}
