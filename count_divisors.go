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
	l, _ := strconv.Atoi(input.Text())
	input.Scan()
	r, _ := strconv.Atoi(input.Text())
	input.Scan()
	k, _ := strconv.Atoi(input.Text())
	fmt.Println((r/k)-((l-1)/k))

}
