package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// var n int
	// fmt.Print("Jumlah Array : ")
	// fmt.Scan(&n)
	// perfectTwo(n)

	// var s string
	// fmt.Print("Input : ")
	// fmt.Scan(&s)
	// programmerIT(s)

	var digit int
	fmt.Print("Input Digit : ")
	fmt.Scan(&digit)
	biggestPalindrome(digit)

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

func programmerIT(s string) {
	dictionary := []string{"pro", "gram", "merit", "program", "it", "programmer"}
	result := []string{}
	for i := 0; i < len(dictionary); i++ {
		if strings.Contains(s, dictionary[i]) {
			result = append(result, dictionary[i])
		}
	}
	if len(result) > 0 {
		fmt.Println(result)
	} else {
		fmt.Println("<no way>")
	}
}

func biggestPalindrome(digit int) {
	result := []int{}
	n := 1
	m := 1
	for {
		if digit == 1 {
			n = 1
			break
		}
		if len(strconv.Itoa(n)) < digit {
			n++
		} else {
			break
		}
	}

	mystart := n
	for {
		if len(strconv.Itoa(m)) > digit {
			m = mystart
			n++
		}
		if len(strconv.Itoa(n)) > digit {
			break
		}

		t := n * m

		if isPalindrome(t) {
			if contains(result, t) == false {
				result = append(result, t)
			}

		}

		m++
	}
	sort.Ints(result)
	fmt.Println(result[len(result)-1])
}

func isPalindrome(number int) bool {
	var remainder, temp int
	var reverse int = 0

	temp = number

	// For Loop used in format of While Loop
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == reverse {
		return true
	} else {
		return false
	}
}

func contains(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
