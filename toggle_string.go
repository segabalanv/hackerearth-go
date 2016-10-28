package main
 
import (
	"fmt"
	"bufio"
	"os"
	"unicode"
)
 
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := []rune(input.Text())
	for i := 0; i < len(str); i++ {
		if unicode.IsLower(str[i]) {
			str[i] = unicode.ToUpper(str[i])
		} else if unicode.IsUpper(str[i]) {
			str[i] = unicode.ToLower(str[i])
		}
	}
	fmt.Println(string(str))
}