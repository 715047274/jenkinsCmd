package testreport

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/smtp"
)

type MailClient struct {
	Domain string
}

func (mc *MailClient) getMXRecord() (string, error) {
	mxs, err := net.LookupMX(mc.Domain)
	if err != nil || len(mxs) == 0 {
		return "", fmt.Errorf("failed to lookup MX record for domain: %s", mc.Domain)
	}
	return mxs[0].Host, nil
}

func (mc *MailClient) SendHTMLEmailWithAttachment(sender, recipient, subject, htmlContent, attachmentPath string) error {
	mx, err := mc.getMXRecord()
	if err != nil {
		return err
	}

	c, err := smtp.Dial(mx + ":25")
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer c.Quit()

	if err := c.Mail(sender); err != nil {
		return err
	}
	if err := c.Rcpt(recipient); err != nil {
		return err
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/mixed; boundary=%s\r\n\r\n", sender, recipient, subject, writer.Boundary())
	buf.WriteString(headers)

	htmlPart, _ := writer.CreatePart(map[string][]string{
		"Content-Type": {"text/html; charset=UTF-8"},
	})
	htmlPart.Write([]byte(htmlContent))

	if attachmentPath != "" {
		fileContent, _ := ioutil.ReadFile(attachmentPath)
		attachmentPart, _ := writer.CreatePart(map[string][]string{
			"Content-Disposition": {fmt.Sprintf("attachment; filename=\"%s\"", attachmentPath)},
		})
		attachmentPart.Write([]byte(base64.StdEncoding.EncodeToString(fileContent)))
	}

	writer.Close()
	wc, _ := c.Data()
	wc.Write(buf.Bytes())
	wc.Close()

	return nil
}
