package word2

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "hello 你好"
	fmt.Println("len(str):", len(str))
	fmt.Println("utf8.RuneCountInString(str):", utf8.RuneCountInString(str))
	fmt.Println("len([]rune(str)):", len([]rune(str)))
}
