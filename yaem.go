package yaem

import (
	"net/smtp"
	"fmt"
)

const (
	// text messaging email gateway formats
	ATT_EMAIL_FORMAT = "%s@txt.att.net"
	METRO_PCS_EMAIL_FORMAT = "%s@mymetropcs.com"
	T_MOBILE_EMAIL_FORMAT = "%s@tmomail.net"
	US_CELLULAR_EMAIL_FORMAT = "%s@email.uscc.net"
	CRICKET_WIRELESS_EMAIL_FORMAT = "%s@sms.mycricket.com"
	SPRINT_EMAIL_FORMAT = "%s@messaging.sprintpcs.com"
	TRACFONE_EMAIL_FORMAT = "%s@mmst5.tracfone.com"
	VERIZON_EMAIL_FORMAT = "%s@vtext.com"

	//smtp hostname and port numbers
	GMAIL_SMTP_HOSTNAME = "smtp.gmail.com"
	GMAIL_SMTP_PORT = 587
	YAHOO_SMTP_HOSTNAME = "smtp.mail.yahoo.com"
	YAHOO_SMTP_PORT = 465
)

type Emailer struct {
	SmtpUsername string
	SmtpPassword string
	SmtpHostname string
	SmtpPort int
}

func New(username string, password string, hostname string, port int) *Emailer {
	e := new (Emailer)
	e.SmtpUsername = username
	e.SmtpPassword = password
	e.SmtpHostname = hostname
	e.SmtpPort = port
	return e
}

func (e *Emailer) SendText(phone_number string, provider_email_format string, subject string, message string) error {
	body := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message)
	auth := smtp.PlainAuth("", e.SmtpUsername, e.SmtpPassword, e.SmtpHostname)
	connection := fmt.Sprintf("%s:%d", e.SmtpHostname, e.SmtpPort)
	to := fmt.Sprintf(provider_email_format, phone_number)
	err := smtp.SendMail(connection, auth, e.SmtpUsername, []string{to}, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func (e *Emailer) SendEmail(to []string, subject string, message string) error {
	body := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message)
	auth := smtp.PlainAuth("", e.SmtpUsername, e.SmtpPassword, e.SmtpHostname)
	connection := fmt.Sprintf("%s:%d", e.SmtpHostname, e.SmtpPort)
	err := smtp.SendMail(connection, auth, e.SmtpUsername, to, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func (e *Emailer) Dump() {
	fmt.Println("Emailer")
	fmt.Println("-------")
	fmt.Println("\tSmtpUsername:          ", e.SmtpUsername)
	fmt.Println("\tSmtpPassword:          ", e.SmtpPassword)
	fmt.Println("\tSmtpHostname:          ", e.SmtpHostname)
	fmt.Println("\tSmtpPort:              ", e.SmtpPort)
}