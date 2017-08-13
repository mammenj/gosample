package main

import (
	"log"
	"os"

	"github.com/mammenj/gosample/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		log.Printf("#%-5d  %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
