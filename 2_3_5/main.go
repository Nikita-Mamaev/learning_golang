package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	string1 := scanner.Text()
	_ = scanner.Scan()
	string2 := scanner.Text()
	_ = scanner.Scan()
	string3 := scanner.Text()
	_ = scanner.Scan()
	string4 := scanner.Text()

	fmt.Println(string2 + string1 + string3 + string1 + string4)
}
