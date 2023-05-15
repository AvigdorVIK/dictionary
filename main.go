package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dictionary := make(map[string]string)

	scaner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("enter the word and its translation separated by a space:")

		scaner.Scan()
		input := scaner.Text()

		if input == "" {
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("incorrectly entered format. enter words separated by spaces.")
			continue
		}

		key := parts[0]
		value := parts[1]

		dictionary[key] = value

	}

	fmt.Println("list of your words")
	for key, value := range dictionary {
		fmt.Printf("%s:%s\n", key, value)
	}

}
