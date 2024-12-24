package mailer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net"
	"net/smtp"
	"os"
	"strings"
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

func (mc *MailClient) SendHTMLEmailWithAttachment(sender, allRecipients, subject, htmlContent, attachmentPath string) error {
	// Resolve the SMTP server
	mx, err := mc.getMXRecord()
	if err != nil {
		return err
	}

	// Parse the recipients into To and CC lists
	to := []string{}
	cc := []string{}
	recipients := strings.Split(allRecipients, ",")
	for _, recipient := range recipients {
		recipient = strings.TrimSpace(recipient)
		if strings.HasPrefix(recipient, "cc_") {
			cc = append(cc, strings.TrimPrefix(recipient, "cc_"))
		} else {
			to = append(to, recipient)
		}
	}

	// Establish SMTP connection
	c, err := smtp.Dial(mx + ":25")
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer c.Quit()

	// Set the sender
	if err := c.Mail(sender); err != nil {
		return err
	}

	// Add all recipients (To + CC) to SMTP RCPT command
	allSMTPRecipients := append(to, cc...)
	for _, recipient := range allSMTPRecipients {
		if err := c.Rcpt(recipient); err != nil {
			return err
		}
	}

	// Construct the email
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Email headers
	headers := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\n", sender, strings.Join(to, ", "),
	)

	if len(cc) > 0 {
		headers += fmt.Sprintf("Cc: %s\r\n", strings.Join(cc, ", "))
	}

	headers += fmt.Sprintf(
		"Subject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/mixed; boundary=%s\r\n\r\n",
		subject, writer.Boundary(),
	)
	buf.WriteString(headers)

	// Add HTML content
	htmlPart, _ := writer.CreatePart(map[string][]string{
		"Content-Type": {"text/html; charset=UTF-8"},
	})
	htmlPart.Write([]byte(htmlContent))

	// Add attachment if specified
	if attachmentPath != "" {
		fileContent, _ := os.ReadFile(attachmentPath)
		attachmentPart, _ := writer.CreatePart(map[string][]string{
			"Content-Disposition": {fmt.Sprintf("attachment; filename=\"%s\"", attachmentPath)},
		})
		attachmentPart.Write([]byte(base64.StdEncoding.EncodeToString(fileContent)))
	}

	// Close the writer
	writer.Close()

	// Send the email data
	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = wc.Write(buf.Bytes())
	if err != nil {
		return err
	}
	wc.Close()

	return nil
}
