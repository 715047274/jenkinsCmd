package main

import (
	"fmt"
	"github.com/nilslice/email"
)

// this is working one
func main() {
	msg := email.Message{
		To:      "k.zhang@ceridian.com",    // do not add < > or name in quotes
		From:    "autotest@yourdomain.com", // do not add < > or name in quotes
		Subject: "A simple email testing 4",
		Body:    "Plain text email body. HTML not yet supported, but send a PR!",
	}
	//550 5.7.0 Email rejected per DMARC policy
	err := msg.Send()
	if err != nil {
		fmt.Println(err)
	}
}
