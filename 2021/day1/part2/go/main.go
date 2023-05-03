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

		// FIX -> It's giving the wrong number of count
		if len(vals) > 2 && sum > vals[len(vals)-2] {
			count++
		}
		if start == end {
			start = end
			end = start + 3

			vals = append(vals, sum)
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
