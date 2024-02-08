package helpers

import (
	"api-artha/src/configs"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

type SRecipient struct {
	To  []string
	Cc  []string
	Bcc []string
}

func SendMail(recipient SRecipient, subject string, bodyContent *strings.Builder) error {
	configs := configs.SMTPConfig()
	auth := smtp.PlainAuth("", configs.Username, configs.Password, configs.Host)
	smtpAddr := fmt.Sprintf("%s:%s", configs.Host, configs.Port)

	tlsConfig := &tls.Config{
		ServerName: configs.Host,
	}

	client, err := smtp.Dial(smtpAddr)
	if err != nil {
		return err
	}

	if err := client.StartTLS(tlsConfig); err != nil {
		return err
	}

	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(configs.Username); err != nil {
		return err
	}

	allRecipients := append(recipient.To, recipient.Cc...)
	allRecipients = append(allRecipients, recipient.Bcc...)
	for _, recipient := range allRecipients {
		if err := client.Rcpt(recipient); err != nil {
			return err
		}
	}

	wc, err := client.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = fmt.Fprintf(wc, "From: Artha <%s>\r\n", configs.Username)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(wc, "To: %s\r\n", strings.Join(recipient.To, ", "))
	if err != nil {
		return err
	}

	if recipient.Cc != nil {
		_, err = fmt.Fprintf(wc, "Cc: %s\r\n", strings.Join(recipient.Cc, ", "))
		if err != nil {
			return err
		}
	}

	if recipient.Bcc != nil {
		_, err = fmt.Fprintf(wc, "Bcc: %s\r\n", strings.Join(recipient.Bcc, ", "))
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(wc, "Subject: %s\r\n", subject)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(wc, "MIME-Version: 1.0\r\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(wc, "Content-Type: text/html; charset=UTF-8\r\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(wc, "\r\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(wc, bodyContent.String())
	if err != nil {
		return err
	}

	return nil
}
