package main

import "fmt"

func main() {
	const lightspeed = 299792
	var distance = 5600000
	fmt.Println(distance/lightspeed, "senconds")
	distance = 401000000
	fmt.Println(distance/lightspeed, "senconds")
}
