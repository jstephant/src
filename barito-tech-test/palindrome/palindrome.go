package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var digit int
	fmt.Print("Input Digit : ")
	fmt.Scan(&digit)
	biggestPalindrome(digit)
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

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
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
