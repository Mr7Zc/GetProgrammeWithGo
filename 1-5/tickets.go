package main

import (
	"fmt"
	"math/rand"
)

const sencondsPerDay = 86400

func main() {
	distance := 62100000
	company := ""
	trip := ""
	fmt.Println("Spaceline      Days Trip type Price")
	fmt.Println("===================================")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "SA"
		case 1:
			company = "SX"
		case 2:
			company = "VG"
		}

		speed := rand.Intn(15) + 16
		duration := distance / speed / sencondsPerDay
		price := 20.0 + speed
		if rand.Intn(2) == 1 {
			trip = "r-t"
			price += 2
		} else {
			trip = "o-d"
		}
		fmt.Printf("%-10v %4v %-10v $%4v\n", company, duration, trip, price)
	}
}
