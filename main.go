package main

import (
	"flag"
	"fmt"
)

func main() {
	var position int
	flag.IntVar(&position, "fibPos", 0, "Integer")
	flag.Parse()

	result := fibonnaci(position)

	fmt.Print(result)
}

func fibonnaci(position int) int {
	if position <= 1 {
		return position
	} else {
		fibSeq := [3]int{0, 1, 1}

		for index := 2; index < position; index++ {
			fmt.Println(fibSeq)

			f1 := fibSeq[1]
			f2 := fibSeq[2]

			fibSeq[0] = f1
			fibSeq[1] = f2
			fibSeq[2] = f1 + f2
		}

		return fibSeq[2]
	}
}
