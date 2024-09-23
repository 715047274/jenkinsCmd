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
	//var (
	//	mx  string
	//	err error
	//)
	//msg := []byte("To: k.zhang@ceridian.com\r\n" +
	//	"Subject: Hi there!\r\n" +
	//	"Content-Type: text/plain; charset=UTF-8\r\n" +
	//	"\r\n" +
	//	"Hi!\r\n")
	//to := make([]string, 1)
	//to[0] = "autotest@dayforce.com"
	//
	//mx, err = getMXRecord("corpadds.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//res := smtp.SendMail(mx+":25",
	//	nil, /* this is the optional Auth */
	//	"autotest@dayforce.com", to, msg)
	//if res != nil {
	//	fmt.Println(res)
	//}
	//log.Fatal(res)

	///////////////////////////////////////

	// Connect to the remote SMTP server.
	//var (
	//	mx  string
	//	err error
	//)
	//mx, err = getMXRecord("corpadds.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c, err := smtp.Dial(mx + ":25")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// Set the sender and recipient.
	//c.Mail("no-reply@yourdomain.com")
	//c.Rcpt("k.zhang@ceridian.com")
	//// Send the email body.
	//wc, err := c.Data()
	//
	//_, err = wc.Write(msg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////How do I get the response here ??
	//err = wc.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err != nil {
	//	log.Fatal(err)
	//
	//}
	//
	//defer wc.Close()
	//buf := bytes.NewBufferString("This is the email body.")
	//if _, err = buf.WriteTo(wc); err != nil {
	//	log.Fatal(err)
	//}

}

//	func getMXRecord(to string) (mx string, err error) {
//		// var e *mail.Address
//		//e, err = mail.ParseAddress(to)
//		//if err != nil {
//		//	return
//		//}
//
//		domain := to
//
//		var mxs []*net.MX
//		mxs, err = net.LookupMX(domain)
//
//		if err != nil {
//			return
//		}
//
//		for _, x := range mxs {
//			mx = x.Host
//			fmt.Println(mx)
//			return
//		}
//
//		return
//	}
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

	if err = c.Mail("no-reply@yourdomain.com"); err != nil {
		log.Fatal(err)
	}

	if err = c.Rcpt("k.zhang@ceridian.com"); err != nil {
		log.Fatal(err)
	}

	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)

	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close() // make sure WriterCloser gets closed

	return
}

func getMXRecord(to string) (mx string, err error) {
	// var e *mail.Address
	//e, err = mail.ParseAddress(to)
	//if err != nil {
	//	return
	//}

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
