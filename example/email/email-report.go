package main

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
)

func main() {
	msg := []byte("Here is a string....")
	var (
		_   int
		_   string
		err error
	)
	_, _, err = sendEmail(msg)
	if err != nil {
		log.Fatal(err)
	}

}

func sendEmail(msg []byte) (code int, message string, err error) {
	var (
		mx string
	)
	mx, err = getMXRecord("corpadds.com")
	if err != nil {
		log.Fatal(err)
	}
	c, err := smtp.Dial(mx + ":25")
	if err != nil {
		log.Fatal(err)

	}
	// Set the sender and recipient.
	defer c.Quit() // make sure to quit the Client

	var sender = "autotest@yourdomain.com"

	if err = c.Mail(sender); err != nil {
		log.Fatal(err)
	}
	if err = c.Rcpt("k.zhang@ceridian.com"); err != nil {
		log.Fatal(err)
	}

	wc, err := c.Data()
	defer wc.Close() // make sure WriterCloser gets closed
	if err != nil {
		log.Fatal(err)

	}
	_, err = fmt.Fprintf(wc, string(msg))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getMXRecord(to string) (mx string, err error) {
	domain := to
	var mxs []*net.MX
	mxs, err = net.LookupMX(domain)
	if err != nil {
		return
	}
	for _, x := range mxs {
		mx = x.Host
		fmt.Println(mx)
		return
	}
	return
}
