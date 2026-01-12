package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(findDivisors(24))
	// fmt.Println(extractDigit(1278))
	// starsTriangle1(5)
	// starsTriangle2(5)
	// pyramidNumber(5)
	// numberSeries1(5)
	// numberSeries2(9)
	// fmt.Println(isPalindrome("Aku Usa"))
	// fmt.Println(reverseWord("ABCD"))
	// fmt.Println(checkBraces("())"))
	// fmt.Println(isNumberPalindrome(1214324))
}

// nomor 1
func findDivisors(n int) string {
	res := ""
	for i := 1; i <= n; i++ {
		if n%i == 0 && n != i {
			res += fmt.Sprintf("%d ", i)
		}
	}
	return res
}

// nomor 2
func extractDigit(n int) string {
	digits := ""
	for n > 0 {
		digits += fmt.Sprintf("%d ", n%10)
		n /= 10
	}
	return digits
}

// nomor 3
func starsTriangle1(n int) {
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < i; j++ {
			s += "  "
		}
		for k := 0; k < n-i; k++ {
			s += "* "
		}
		fmt.Println(s)
	}
}

func starsTriangle2(n int) {
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < n-i-1; j++ {
			s += "  "
		}
		for k := 0; k < i+1; k++ {
			s += "* "
		}
		fmt.Println(s)
	}
}

// nomor 4
func pyramidNumber(n int) {
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < n-i-1; j++ {
			s += fmt.Sprintf("%d ", n-j-i)
		}
		for k := 0; k < n-i; k++ {
			s += fmt.Sprintf("%d ", k+1)
		}
		fmt.Println(s)
	}
}

// nomor 5
func numberSeries1(n int) {
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < n; j++ {
			if j%2 == 0 {
				s += fmt.Sprintf("%d ", i+1)
			} else {
				s += fmt.Sprintf("%d ", n-i)
			}
		}
		fmt.Println(s)
	}
}

// nomor 6
func numberSeries2(n int) {
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < n; j++ {
			if (i+j)%2 == 0 {
				s += fmt.Sprintf("%s ", "-")
			} else {
				s += fmt.Sprintf("%d ", j+1)
			}
		}
		fmt.Println(s)
	}
}

// nomor 8
func isPalindrome(word string) bool {
	lowerWord := strings.ToLower(word)
	for i := 0; i < len(lowerWord); i++ {
		if lowerWord[i] != lowerWord[len(lowerWord)-i-1] {
			return false
		}
	}
	return true
}

// nomor 9
func reverseWord(word string) string {
	res := ""
	for i := 0; i < len(word); i++ {
		res += string(word[len(word)-i-1])
	}
	return res
}

// nomor 10
func checkBraces(braces string) bool {
	count := 0
	for i := 0; i < len(braces); i++ {
		if braces[i] == '(' {
			count++
		} else if braces[i] == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

// nomor 11
func isNumberPalindrome(n int) bool {
	str := fmt.Sprintf("%d", n)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
