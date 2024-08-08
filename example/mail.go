package main

//func sendEmail(msg []byte) (code int, message string, err error) {
//	var (
//		mx string
//	)
//	mx, err = getMXRecord("corpadds.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//	c, err := smtp.Dial(mx + ":25")
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	// Set the sender and recipient.
//	c.Mail("no-reply@yourdomain.com")
//	c.Rcpt("k.zhang@ceridian.com")
//
//	defer c.Quit() // make sure to quit the Client
//
//	if err = c.Mail("no-reply@yourdomain.com"); err != nil {
//		return
//	}
//
//	if err = c.Rcpt("k.zhang@ceridian.com"); err != nil {
//		return
//	}
//
//	wc, err := c.Data()
//	if err != nil {
//		return
//	}
//	defer wc.Close() // make sure WriterCloser gets closed
//
//	_, err = wc.Write(msg)
//	if err != nil {
//		return
//	}
//
//	code, message, err = c.Text.ReadResponse(0)
//	return
//}
//
//func getMXRecord(to string) (mx string, err error) {
//	// var e *mail.Address
//	//e, err = mail.ParseAddress(to)
//	//if err != nil {
//	//	return
//	//}
//
//	domain := to
//
//	var mxs []*net.MX
//	mxs, err = net.LookupMX(domain)
//
//	if err != nil {
//		return
//	}
//
//	for _, x := range mxs {
//		mx = x.Host
//		fmt.Println(mx)
//		return
//	}
//
//	return
//}
