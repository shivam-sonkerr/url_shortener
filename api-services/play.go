package main

import (
	"fmt"
	"log"
	"net/url"
	_ "net/url"
)

//func urlGenerators() {
//	random := rand.Int()
//	fmt.Println("Random Number is: ", random)
//}

func urlValidator() {

	u, err := url.ParseRequestURI("123")

	fmt.Printf("URL Received: ", u)
	if err != nil {
		log.Fatal("ERROR: ", err, u)
	}

}

//func urlGeneratorsa() string {
//
//	random := rand.IntN(9000) + 1000
//	prefix := "simple-url"
//	fmt.Println(random)
//
//	shortURL := fmt.Sprintf("%s-%d", prefix, random)
//
//	return shortURL
//}

func main() {
	//for i := 3; i < 80; i++ {
	//	fmt.Println("Random number: ", rand.IntN(500))
	//}
	//urlGenerators()
	//urlGeneratorsa()

	urlValidator()
}
