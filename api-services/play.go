package main

import (
	"fmt"
	"math/rand/v2"
)

func urlGenerators() {
	random := rand.Int()
	fmt.Println("Random Number is: ", random)
}

func urlGeneratorsa() string {

	random := rand.IntN(9000) + 1000
	prefix := "simple-url"
	fmt.Println(random)

	shortURL := fmt.Sprintf("%s-%d", prefix, random)

	return shortURL
}

func main() {
	for i := 3; i < 80; i++ {
		fmt.Println("Random number: ", rand.IntN(500))
	}
	urlGenerators()
	urlGeneratorsa()
}
