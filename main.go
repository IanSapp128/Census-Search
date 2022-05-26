package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/pkg/browser"
)

func main() {
	// Set the console window title based on Windows or Linux
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

	var url string
	var table string
	fmt.Println("Enter Table Name: ")
	fmt.Scanln(&table)
	var firstLetter string = string(table[0])

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
		fmt.Scanln()
	}

}
