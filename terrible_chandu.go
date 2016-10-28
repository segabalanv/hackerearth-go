package main
 
import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)
 
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	for i := 0; i < n; i++ {
		input.Scan()
		str := []rune(input.Text())
		for j, k := 0, len(str)-1; j < k; j, k = j+1, k-1 {
			str[j], str[k] = str[k], str[j]
		}
		fmt.Println(string(str))
	}
}