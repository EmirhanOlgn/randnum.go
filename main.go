package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var min, max, randnum int

	for {
		fmt.Print("Enter the min value of the random number: ")
		fmt.Scan(&min)
		fmt.Print("Enter the min value of the random number: ")
		fmt.Scan(&max)

		if max-min < 15 {
			fmt.Println("Minimum and maximum values must have a difference of at least 15")
		} else {
			break
		}
	}

	randnum = randomInteger(min, max)
	forecast := 1

	for {
		var input, difference int
		fmt.Print("Enter a number: ")
		fmt.Scan(&input)

		if input < min || input > max {
			fmt.Println("Please do not exceed the minimum and maximum limits")
		} else {
			difference = int(math.Abs(float64(input - randnum)))

			if input == randnum {
				fmt.Printf("You guessed correctly! You got the correct guess after %v attempts\n", forecast)
				break
			}

			forecast++

			if difference <= 5 {
				fmt.Println("You're so close!")
			} else if difference <= 10 {
				fmt.Println("You are close..")
			} else {
				fmt.Println("You are far away...")
			}
		}
	}
}

func randomInteger(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}
