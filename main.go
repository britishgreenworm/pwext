package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args[1:]

	if args[0] == "/help" || args[0] == "--help" || args[0] == "-h" || args[0] == "-help" {
		fmt.Println("\nExtracts passwords from a file.")
		fmt.Println()
		fmt.Println("usage: pwext {file location} {password length} {exluded characters} {optional: ignore prefix length}")
		return
	}

	if len(args) < 3 {
		fmt.Println("Arguments are required, usage: pwext {file location} {password length} {exluded characters} {optional: ignore prefix length}")
		return
	}

	fileLocation := args[0]
	passwordLength, _ := strconv.Atoi(args[1])
	excludeRunes := args[2]
	ignorePrefix := 0

	if len(args) > 3 {
		ignorePrefix, _ = strconv.Atoi(args[3])
	}

	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for line := 1; scanner.Scan(); line++ {
		runes := []rune(scanner.Text())

		for x := ignorePrefix; x <= (len(scanner.Text()) - passwordLength); {
			if !strings.ContainsAny(string(runes[x:(x+passwordLength)]), excludeRunes) {
				fmt.Println(string(runes[x : x+passwordLength]))
				x += passwordLength
			} else {
				x++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
