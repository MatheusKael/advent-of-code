package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("../../input.txt")

	if err != nil {
		panic(err)
	}

	s := string(file)

	vals := strings.Fields(s)

	count := 0

	prev, err := strconv.Atoi(vals[0])

	if err != nil {
		panic(err)
	}

	for _, val := range vals {

		n, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		if prev < n {
			prev = n
			count++
		} else if prev > n {
			prev = n
		}

	}

	fmt.Println(count)

}
