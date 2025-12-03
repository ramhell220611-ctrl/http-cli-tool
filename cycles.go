package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	adress := flag.String("adress", "", "put adress here")             //flag with url
	data := flag.String("data", "", "JSON data for POST/PUT requests") // flag with JSON
	flag.Parse()

	if len(flag.Args()) > 0 { // This checks your entry
		switch flag.Arg(0) {
		case "get":
			if *adress == "" {
				fmt.Println("You need to put right address")
				os.Exit(1)
			}
			getMethod(*adress)
		case "post":
			if *adress == "" {
				fmt.Println("You need to put right address")
				os.Exit(1)
			}
			postMethod(*adress, *data)
		case "put":
			if *adress == "" {
				fmt.Println("You need to put right address")
				os.Exit(1)
			}
			putMethod(*adress, *data)
		case "del":
			if *adress == "" {
				fmt.Println("You need to put right address")
				os.Exit(1)
			}
			deleteMethod(*adress)
		}
	}
	helpfilFunction()
}
func getMethod(adress string) { // function about get
	cmd := exec.Command("curl", adress)
	output, err := cmd.Output()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(output))
	fmt.Println(adress)
}
func postMethod(adress, data string) { // about post
	cmd := exec.Command("curl", "-s", "-X", "POST",
		"-H", "Content-Type: application/json",
		"-d", data,
		adress)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing POST request: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
	fmt.Println(adress)
}
func putMethod(adress, data string) { // put method
	cmd := exec.Command("curl", "-s", "-X", "PUT",
		"-H", "Content-Type: application/json",
		"-d", data,
		adress)

	output, err := cmd.Output()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(output))
	fmt.Println(adress)
}
func deleteMethod(adress string) { // delete mrthod
	cmd := exec.Command("curl", "-s", "-X", "DELETE", adress)
	output, err := cmd.Output()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(output))
	fmt.Println(adress)
}
func helpfilFunction() { // this calls when you do nothing
	fmt.Println("Oh might be you dont understand how it works, okay no problem")
	fmt.Println("You can use theese methods:")
	fmt.Println("get, post, put, del")
	fmt.Println("But before you firstly need to fill the utilities")
	fmt.Println("Example: ./cycles(or your name of programm) -adress='https://example.org' -data='Your JSON application' get(but it could be all methods)")
}
