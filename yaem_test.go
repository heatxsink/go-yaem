package yaem

import (
	"fmt"
	"testing"
)

var (
	// sorry guy i don't wanna put my personal infos here i promise this works tho!
	username = "your imap username here"
	password = "your imap password here"
	phone_number = "your phone number here"
)

func TestSendEmail(t *testing.T) {
	fmt.Println("yaem.SendEmail")
	e := New(username, password, GMAIL_SMTP_HOSTNAME, GMAIL_SMTP_PORT)
	e.Dump()
	email_address := fmt.Sprintf(ATT_EMAIL_FORMAT, phone_number)
	to := []string{ email_address }
	err := e.SendEmail(to, "Test Send Email", "Body Goes Here")
	if err != nil {
		t.Errorf("ERROR: ", err)
	}
}

func TestSendText(t *testing.T) {
	fmt.Println("yaem.SendText")
	e := New(username, password, GMAIL_SMTP_HOSTNAME, GMAIL_SMTP_PORT)
	e.Dump()
	err := e.SendText(phone_number, ATT_EMAIL_FORMAT, "Test Send Text", "Body Goes Here")
	if err != nil {
		t.Errorf("ERROR: ", err)
	}
}