package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	title := scanner.Text()
	fmt.Println(title, "- лучшая книга!")
}
