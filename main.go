package main

import (
	"flag"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "", "the email to scan")
	flag.Parse()
	//fmt.Println(folder, email)
	if email == "" {
		flag.Usage()
		return
	}

	if folder != "" {
		scan(folder)
		return
	}


	stats(email)
}