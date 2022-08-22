package main

import "fmt"

func power_of_two(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && n&(-1^(n-1)^n) == 0
}

func main() {
	fmt.Println(power_of_two(1 << 28))
}
