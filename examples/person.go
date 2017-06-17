package main

import (
	"fmt"
	"log"

	"github.com/Abramovic/fullcontact-go"
)

func main() {
	client, err := fullcontact.NewClient("")
	fatalOnError(err)

	// Webhook is optional.
	webhook, err := fullcontact.NewWebhook("http://your-web-hook.url/", "ID-123")
	fatalOnError(err)

	// Look up bart@fullcontact.com by email address
	result, err := client.Person.Email("bart@fullcontact.com", nil)
	fatalOnError(err)

	fmt.Println(result)
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
