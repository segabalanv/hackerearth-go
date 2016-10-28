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
	t, _ := strconv.Atoi(input.Text())
	for i := 0; i < t; i++ {
		input.Scan()
		str := []rune(input.Text())
		for j := 0; j < len(str)-1; j++ {
			if str[j] == str[j+1] {
				str[j] = -1
			}
		}
		for j := 0; j < len(str) ; j++ {
			if str[j] != -1 {
				fmt.Printf("%c", str[j])
			}
		}
		fmt.Println()
	}
}