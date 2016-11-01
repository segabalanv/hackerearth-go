package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		input.Scan()
		numbers[i], _ = strconv.Atoi(input.Text())
	}
	print_value := 0
	for i := 0; i < n; i += 3 {
		print_value += numbers[i]
	}
	fmt.Print(print_value, " ")
	print_value = 0
	for i := 1; i < n; i += 3 {
		print_value += numbers[i]
	}
	fmt.Print(print_value, " ")
	print_value = 0
	for i := 2; i < n; i += 3 {
		print_value += numbers[i]
	}
	fmt.Print(print_value)
}
