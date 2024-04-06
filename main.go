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

	maxUint64SupportedPosition := 92

	start := time.Now()

	if position <= maxUint64SupportedPosition {
		result := fibonacci(position)
		fmt.Print("RESULT: ", result)
		fmt.Print("\n\nDigits: ", len(fmt.Sprint(result)))
	} else {
		result := fibonacciBigInt(position)
		fmt.Print("RESULT: ", result)
		fmt.Print("\n\nDigits: ", len(result.String()))
	}

	duration := time.Since(start)

	fmt.Printf("\nPosition: %v\n", position)
	fmt.Printf("Time taken: %v\n", duration)
}

func fibonacci(position int) uint64 {
	if position <= 1 {
		return uint64(position)
	} else {
		fibSeq := [3]uint64{0, 1, 1}

		for index := 2; index < position; index++ {
			fibSeq[0], fibSeq[1] = fibSeq[1], fibSeq[2]

			fibSeq[2] = fibSeq[1] + fibSeq[0]
		}

		return fibSeq[2]
	}
}

func fibonacciBigInt(position int) *big.Int {
	fibSeq := [3]*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1)}

	for index := 2; index < position; index++ {
		fibSeq[0], fibSeq[1] = fibSeq[1], fibSeq[2]

		fibSeq[2] = big.NewInt(0).Add(fibSeq[1], fibSeq[0])
	}

	return big.NewInt(0).Set(fibSeq[2])
}
