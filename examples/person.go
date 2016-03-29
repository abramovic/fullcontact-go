package main

import (
	"fmt"
	"github.com/Abramovic/fullcontact-go"
	"log"
)

func main() {
	client, err := fullcontact.NewClient("my-api-key")
	fatalOnError(err)

	// Webhook is optional.
	webhook, err := fullcontact.NewWebhook("http://your-web-hook.url/", "ID-123")
	fatalOnError(err)

	// Look up john@example.com by email address
	result, err := client.Person.Email("john@example.com", webhook)
	fatalOnError(err)

	fmt.Println(result)
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
