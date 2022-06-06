package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	toSay := strings.Join(os.Args[1:], " ")

	printDialogBox(toSay)
	printCow()
}

func printHelp() {
	fmt.Println("cowsay-go is program for cow say anything in your terminal!")
	fmt.Println("Type: cowsay-go <to say>")
	fmt.Println("For example: cowsay-go hello world")
}

func printDialogBox(toSay string) {
	qtyCarac := len(toSay)

	border := strings.Repeat("-", qtyCarac)

	fmt.Println(border)
	fmt.Println(toSay)
	fmt.Println(border)
}

func printCow() {
	fmt.Println(`
    \
     \
      ------
     / 0 0 /
    / ._. / \____________
    ------\              \
          | _  ______     \
          || ||      ||w|| \
          || ||      || ||
	`)
}
