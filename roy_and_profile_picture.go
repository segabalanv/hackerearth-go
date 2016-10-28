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
	n, _ := strconv.Atoi(input.Text())
	for ; n > 0; n-- {
		input.Scan()
		w, _ := strconv.Atoi(input.Text())
		input.Scan()
		h, _ := strconv.Atoi(input.Text())
		if w < l || h < l {
			fmt.Println("UPLOAD ANOTHER")
		} else if w != h {
			fmt.Println("CROP IT")
		} else {
			fmt.Println("ACCEPTED")
		}
	}
}