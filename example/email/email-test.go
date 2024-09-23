package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

// telnet gmail-smtp-in.l.google.com 25
// openssl s_client -connect gmail-smtp-in.l.google.com:25
// https://gist.github.com/xeoncross/9a5ffa2e4edc2be7681a41f57d1e5c51
func main() {

	var (
		mx  string
		err error
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	from := "autotest@dayforce.com"
	to := "k.zhang@dayforce.com"
	subject := "Test Email from golang"
	body := "This is the email body at " + time.Now().String()

	msg := composeMimeMail(to, from, subject, body)

	mx, err = getMXRecordaa(to)
	if err != nil {
		log.Fatal(err)
	}

	err = smtp.SendMail(mx+":25", nil, from, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func getMXRecordaa(to string) (mx string, err error) {
	var e *mail.Address
	e, err = mail.ParseAddress(to)
	if err != nil {
		return
	}

	domain := strings.Split(e.Address, "@")[1]

	var mxs []*net.MX
	mxs, err = net.LookupMX(domain)

	if err != nil {
		return
	}

	for _, x := range mxs {
		mx = x.Host
		println(mx)
		return
	}

	return
}

// Never fails, tries to format the address if possible
func formatEmailAddress(addr string) string {
	e, err := mail.ParseAddress(addr)
	if err != nil {
		return addr
	}
	return e.String()
}

func encodeRFC2047(str string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Address: str}
	return strings.Trim(addr.String(), " <>")
}

func composeMimeMail(to string, from string, subject string, body string) []byte {
	header := make(map[string]string)
	header["From"] = formatEmailAddress(from)
	header["To"] = formatEmailAddress(to)
	header["Subject"] = encodeRFC2047(subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return []byte(message)
}
