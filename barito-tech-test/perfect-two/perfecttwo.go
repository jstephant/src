package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Jumlah Array : ")
	fmt.Scan(&n)
	perfectTwo(n)
}

func perfectTwo(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Angka %d : ", i)
		fmt.Scan(&a[i])
	}

	var target int
	fmt.Printf("Target : ")
	fmt.Scan(&target)

	output := "<no way>"
	for i := 0; i < len(a); i++ {
		if i+1 < len(a) {
			for j := i + 1; j < len(a); j++ {
				if a[i]+a[j] == target {
					output = fmt.Sprintf("Result : [%d, %d]\n", i, j)
					break
				}
			}
		}
	}
	fmt.Println(output)
}
