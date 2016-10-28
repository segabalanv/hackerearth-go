package main
 
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)
 
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	input.Scan()
	t, _ := strconv.Atoi(input.Text())
	for t > 0 {
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		x, _ := strconv.Atoi(input.Text())
		friends := make([]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			friends[i], _ = strconv.Atoi(input.Text())
		}
		result := "NO"
		for i := 0; i < n; i++ {
			sum := 0
			for j := i; j < n; j++ {
				sum += friends[j]
				if sum > x {
					break
				}
				if sum == x {
					break
				}
			}
			if sum == x {
				result = "YES"
				break
			}
		}
		fmt.Println(result)
		t--
	}
}