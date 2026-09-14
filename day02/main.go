package main

import "fmt"

func sum(numbers []int) int {
	sums := 0
	for _, number := range numbers {
		sums += number
	}
	return sums
}

func max(numbers []int) (int, bool) {
	maximum := 0
	if len(numbers) == 0 {
		return maximum, false
	}

	maximum = numbers[0]
	for _, number := range numbers {
		if maximum < number {
			maximum = number
		}
	}
	return maximum, true
}

func analyze(numbers []int) (minNumber int, maxNumber int, avgNumber float64) {
	fullNumbers := 0
	maxNumber, ok := max(numbers)
	if !ok {
		return 0, 0, 0
	}
	minNumber = numbers[0]
	for _, number := range numbers {
		if minNumber > number {
			minNumber = number
		}
		fullNumbers += number
	}
	avgNumber = float64(fullNumbers) / float64(len(numbers))
	return minNumber, maxNumber, avgNumber
}

func countNumbers(numbers []int) map[int]int {
	countMapNumbers := make(map[int]int)

	for _, number := range numbers {
		countMapNumbers[number] += 1
	}
	return countMapNumbers
}

func unique(numbers []int) []int {
	seen := make(map[int]bool)
	uniqueNumbers := make([]int, 0, len(numbers))

	for _, number := range numbers {
		if !seen[number] {
			uniqueNumbers = append(uniqueNumbers, number)
			seen[number] = true
		}
	}
	return uniqueNumbers
}

func runeCount(text string) int {
	return len([]rune(text))
}

func firstRune(text string) (rune, bool) {
	if len(text) == 0 {
		return 0, false
	}
	runes := []rune(text)
	return runes[0], true
}

func reverseString(text string) string {
	runes := []rune(text)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	//numbers := []int{1, 1, 2, 3, 2, 2, 4, 2}
	//fmt.Println(sum(numbers))
	//fmt.Println(max(numbers))
	//fmt.Println(analyze(numbers))
	//fmt.Println(countNumbers(numbers))
	//fmt.Println(unique(numbers))
	fmt.Println(runeCount("hello"))
	fmt.Println(runeCount("Привет"))
	fmt.Println(runeCount(""))
	fmt.Println(firstRune("Привет"))
	fmt.Println(firstRune("Go"))
	fmt.Println(firstRune(""))
	fmt.Println(reverseString("Привет"))
	fmt.Println(reverseString("Go"))
	fmt.Println(reverseString(""))

}
