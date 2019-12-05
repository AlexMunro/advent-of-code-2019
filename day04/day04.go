package main

import(
	"../utils"
	"path/filepath"
	"fmt"
	"strconv"
)

func validPassword(password int) bool {
	adjacentDigits := false
	for password > 10 { // this will break if your lower bound starts with a 0, soz
		lastDigit := password % 10
		password /= 10
		if lastDigit < password % 10 {
			return false
		}
		if lastDigit == password % 10 {
			adjacentDigits = true
		}
	}
	return adjacentDigits
}

func extraValidPassword(password int) bool {
	// separating out non-decreasing and adjacent concerns
	passwordCopy := password
	for passwordCopy > 10 { // still gonna break if your lower bound starts with a 0
		lastDigit := passwordCopy % 10
		passwordCopy /= 10
		if lastDigit < passwordCopy % 10 {
			return false
		}
	}

	for password > 10 {
		lastDigit := password % 10
		password /= 10
		if lastDigit == password % 10 {
			if password < 10 || (password / 10) % 10 != lastDigit {
				return true
			}
			password /= 10
			for (lastDigit == password % 10 && password > 10){
				password /= 10
			}
		}
	}
	return false
}

func passwordsInRange(low, high int, criteria func(int) bool) int {
	count := 0
	for i := low; i <= high; i++ {
		if criteria(i){
			count++
		}
	}
	return count
}

func main() {
	filename, _ := filepath.Abs("./input.txt")
	input := utils.GetInputSingleString(filename)
	low, _ := strconv.Atoi(input[:6])
	high, _ := strconv.Atoi(input[7:])
	
	validPasswordCount := passwordsInRange(low, high, validPassword)
	fmt.Printf("The answer to part one is %d\n", validPasswordCount)

	extraValidPasswordCount := passwordsInRange(low, high, extraValidPassword)
	fmt.Printf("The answer to part two is %d\n", extraValidPasswordCount)
}