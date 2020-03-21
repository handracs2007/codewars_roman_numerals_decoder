package main

import "fmt"

func Decode(roman string) int {
	var romanMap = make(map[rune]int)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	var result = 0
	var prev = -1

	for _, val := range roman {
		currentVal := romanMap[val]
		result += currentVal

		if prev != -1 && currentVal > prev {
			result -= currentVal
			result -= prev
			result += currentVal - prev
		}

		prev = currentVal
	}

	return result
}

func main() {
	fmt.Println(Decode("MCMXC"))
	fmt.Println(Decode("MMVIII"))
	fmt.Println(Decode("MDCLXVI"))
}
