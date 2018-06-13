package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Just Main Hanzo")
	fmt.Println("Import data [ MongoDB ] ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Name :" + text)
	fmt.Print("Enter String: ")
	text2, _ := reader.ReadString('\n')
	fmt.Println("Name2 :" + text2)
}
