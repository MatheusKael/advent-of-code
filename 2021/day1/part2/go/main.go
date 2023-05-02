package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sol()

}
func sol2() {

	data, err := os.ReadFile("../../input.txt")

	if err != nil {
		panic(err)
	}

	stringFirst := string(data)

	s := strings.Fields(stringFirst)

	vals := make([]interface{}, 0)
	end := 3
	sum := 0
	count := 0

	for idx := range s {
		for i := idx; i < idx+end && idx+end < len(s)+1; i++ {

			number, err := strconv.Atoi(s[i])
			fmt.Println(number)
			if err != nil {
				panic(err)
			}
			sum += number

		}

		if idx != 0 {

			n, err := strconv.Atoi(s[idx-1])
			if err != nil {
				panic(err)
			}
			if sum > n {
				count++
			}
		}
		vals = append(vals, sum)
		sum = 0
	}
	fmt.Println(count)

}

func sol() {

	data, err := os.ReadFile("../../input.txt")

	if err != nil {
		panic(err)
	}

	stringFirst := string(data)

	s := strings.Fields(stringFirst)

	vals := []int{}
	start := 0
	end := start + 3
	sum := 0
	count := 0

	length := len(s)

	for start != length {

		if start == end {
			start = end
			end = start + 3

			vals = append(vals, sum)
			// FIX -> Sum needs to be bigger than the previous sum value in vals
			if sum > vals[len(vals)-2] && len(vals) > 2 {
				fmt.Println(vals[len(vals)-1])
			}
			sum = 0
		}

		n, err := strconv.Atoi(s[start])

		if err != nil {
			panic(err)
		}

		sum += n

		start++

	}
	fmt.Println(vals)
	fmt.Println(count)
}
