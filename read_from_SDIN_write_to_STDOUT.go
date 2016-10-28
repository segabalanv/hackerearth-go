package main
 
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
 
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	input.Scan()
	fmt.Println(2*n)
	fmt.Println(input.Text())
}