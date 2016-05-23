package main
import (
	"fmt"
	"strings"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var i int
    fmt.Scan(&i)
    for temp := 0; temp < i; temp++ {
        fmt.Println(leftPad(strings.Repeat("#", temp + 1), i))
    }

}

func leftPad(s string, n int) string {
    return strings.Repeat(" ", n-len(s)) + s
}
