package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	functinonal := make(map[string]string)

	for true {
		fmt.Println("Salam! I'm Redis, SoftClub edition :)")
		fmt.Println("I can store any data by a specific key, like Redis. " +
			"Unfortunately for now, I can store any data in string format. " +
			"But I'm gonna be better :)")

		fmt.Println()
		fmt.Println("Here are the list of my abilities:")

		fmt.Println("1. Add a value under a specified key (add)")
		fmt.Println("2. Get a value of a specified key (get)")
		fmt.Println("3. Delete a value under a specified key (delete)")
		fmt.Println("4. Get all the key-value pairs (get-all)")
		fmt.Println("5. Remove all the key-value pairs from storage (clear)")
		fmt.Println("6. Quit (quit)")

		var Scaner string
		fmt.Println("\nType a command, you want to run:")
		fmt.Scan(&Scaner)

		if Scaner == "quit" {
			fmt.Println("Goodbye!")
			return
		}

		if Scaner == "add" {
			stdin := bufio.NewScanner(os.Stdin)

			fmt.Println("Okay, let's add something to storage.")
			fmt.Println("Enter a key:")
			stdin.Scan()
			addkey := stdin.Text()

			fmt.Println("Good. Now enter the value you want to save:")
			stdin.Scan()
			addvalue := stdin.Text()

			fmt.Println("Nice. I've added this data to storage. You can check it by getting your key.")
			functinonal[addkey] = addvalue
		}

		if Scaner == "get" {
			stdin := bufio.NewScanner(os.Stdin)
			fmt.Println("No problem, enter the key, you want to get from storage:")
			stdin.Scan()
			addkey := stdin.Text()

			_, exists := functinonal[addkey]

			if !exists {
				fmt.Println("Oh, seems like there is no such key in storage. Sorry.")
			} else {
				fmt.Println("Here it is:")
				fmt.Println(functinonal[addkey])
			}

		}
		if Scaner == "delete" {
			stdin := bufio.NewScanner(os.Stdin)
			fmt.Println("Okay, I'll delete what you want, but be careful.")
			fmt.Println("Enter a key, you want to delete:")
			stdin.Scan()
			addkey := stdin.Text()
			delete(functinonal, addkey)
			fmt.Println("Done! I've deleted it.")
		}

		if Scaner == "get-all" {
			fmt.Println("Here is the list of all the key-value pairs in the storage:")
			fmt.Println(functinonal)
		}

		if Scaner == "clear" {
			fmt.Println("Removing all the data from the storage.")
			functinonal = make(map[string]string)

			fmt.Println("Done! Storage is cleared!")
		}

		if Scaner == "" {
			fmt.Println("Bro, I have no this command on the list!")
		}

	}
}
