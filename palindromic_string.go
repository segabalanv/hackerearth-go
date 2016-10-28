package main
 
import (
	"fmt"
	"bufio"
	"os"
)
 
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	output := "YES"
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		if (str[i] != str[j]) {
			output = "NO"
			break;
		}
	}
	fmt.Println(output)
}