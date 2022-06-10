package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(3)

	messages := make(chan string)
	var dataName []string

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		dataName = append(dataName, <-messages)
	}
	fmt.Println(dataName)
}
