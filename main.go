package main

import "fmt"

func main() {
	userData, err := ParseCmd()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(userData.ToJSON())
}
