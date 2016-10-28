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
	a := make([]int, n)
	for i:=0 ; i < n; i++ {
		input.Scan()
		a[i], _ = strconv.Atoi(input.Text())
	}
	answer := 1
	for i:=0 ; i < n; i++ {
		answer = (answer * a[i]) % 1000000007
	}
	fmt.Println(answer)
}