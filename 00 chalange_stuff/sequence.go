package main

import "fmt"

func main() {
	maxLengt := 1000000

	fmt.Print("enter your number for maxLengt (by default it's 1,000,000)")
	fmt.Scan(&maxLengt)
	if maxLengt == 0 {
		maxLengt = 1000000
	}

	tempLen := 0
	res := 0

	for i := 1; i <= maxLengt; i++ {
		var sequence []int

		sequence = append(sequence, i)

		for true {
			lastNode := func() int {
				return sequence[len(sequence)-1]
			}

			nextNode := func(n int) int {
				if n%2 == 0 {
					return n / 2
				} else {
					return (3 * n) + 1
				}
			}

			sequence = append(sequence, nextNode(lastNode()))

			if lastNode() == 1 {
				break
			}
		}

		if tempLen < len(sequence) {
			tempLen = len(sequence)
			res = i
		}
	}

	fmt.Print(res)
}

/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Which starting number, under one million, produces the longest chain?
*/
