package main
 
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
 
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		input.Scan()
		a[i], _ = strconv.Atoi(input.Text())
	}
	for i := 0; i < n; i++ {
		input.Scan()
		b[i], _ = strconv.Atoi(input.Text())
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i]+b[i])
	}
}