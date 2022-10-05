package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/pkg/browser"
)

func main() {
	var firstArg string

	// Check if there is at least 1 arg passed into the script
	// if so, firstArg will not be equal to "" (empty)
	if len(os.Args) > 1 {
		firstArg = os.Args[1]
	} else {

		// Set the console window title based on Windows or *NIX
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command("cmd", "/C", "title Census Search")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			fmt.Print("\033]0;Census Search\007")
		}
	}

	var url string
	var table string
	var firstLetter string

	/* If firstArg == "" then we can ask the user for their input
	but if it does have a value, just set their value to the firstArg and show results
	that way it can be run from a command prompt with just the argument passed.*/
	if firstArg == "" {
		fmt.Println("Enter Table Name: ")
		fmt.Scanln(&table)
		firstLetter = string(table[0])
	} else {
		firstLetter = string(firstArg[0])
		table = firstArg
	}

	switch {
	case firstLetter == "B":
		url = "https://api.census.gov/data/2020/acs/acs5/groups/" + table + ".html"
		browser.OpenURL(url)
	case firstLetter == "S":
		url = "https://api.census.gov/data/2020/acs/acs5/subject/groups/" + table + ".html"
		browser.OpenURL(url)
	case firstLetter == "D":
		url = "https://api.census.gov/data/2020/acs/acs5/profile/groups/" + table + ".html"
		browser.OpenURL(url)
	default:
		fmt.Println("Not supported. Table name must begin with B, D, or S.")
		// Wait for user input if they're not running directly from command line
		if firstArg == "" {
			fmt.Scanln()
		}

	}

}
