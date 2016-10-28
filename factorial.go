package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
 
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	fmt.Println(fact)
}