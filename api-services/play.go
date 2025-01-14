package main

import (
	"fmt"
	"math/rand/v2"
)

func urlGenerators() {
	random := rand.Int()
	fmt.Println("Random Number is: ", random)
}

func urlGeneratorsa() {
	random := rand.Int()

	prefix := "simple-url"
	fmt.Println(random)

	concat := fmt.Sprintf("%s-%d", prefix, random)

	fmt.Println(concat)
}

func main() {
	for i := 3; i < 80; i++ {
		fmt.Println("Random number: ", rand.IntN(500))
	}
	urlGenerators()
	urlGeneratorsa()
}
