package main

import (
	"flag"
	"fmt"
	"math/big"
	"time"
)

func main() {
	var position int
	flag.IntVar(&position, "fibPos", 0, "Integer")
	flag.Parse()

	if position < 0 {
		fmt.Println("Please provide a non-negative integer for the Fibonacci position.")
		return
	}

	start := time.Now()
	result := fibonacci(position)
	duration := time.Since(start)
	
	fmt.Print("RESULT: ", result)
	fmt.Print("\n\nDigits: ", len(fmt.Sprint(result)))
	fmt.Printf("\nPosition: %v\n", position)
	fmt.Printf("Time taken: %v\n", duration)
}

func fibonacci(position int) *big.Int {
	if position <= 1 {
		return big.NewInt(int64(position))
	} else {
		fibSeq := [3]*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1)}

		for index := 2; index < position; index++ {
			fibSeq[0], fibSeq[1] = fibSeq[1], fibSeq[2]
			fibSeq[2] = big.NewInt(0).Add(fibSeq[1], fibSeq[0])
		}

		return big.NewInt(0).Set(fibSeq[2])
	}
}
