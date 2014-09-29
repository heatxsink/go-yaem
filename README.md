go-yaem
=======

I know I know. Yet another email abstraction or 'yeam'! This is based around sending sms messages via mobile carrier sms email gateways via some other imap based email account. But you can also just send email too. I use this in my projects when I need to send a txt message to my phone (think ghetto alerting on various things ;-P).

Usage
-----

Please take a look at the tests! Below is a quick example of how you could use the yaem package.

```go

import (
	"fmt"
	"github.com/heatxsink/yaem"
)

func main() {
	e := yaem.New("your imap username", "your imap password", yaem.GMAIL_SMTP_HOSTNAME, yaem.GMAIL_SMTP_PORT)
	// send a text message
	phone_number := "5555555555"
	err := e.SendText(phone_number, yaem.ATT_EMAIL_FORMAT, "Test Send Text", "Body Goes Here")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	// or send da emailz
	reciepents := []string{ "someone1@gmail.com", "someone2@gmail.com" }
	err := e.SendEmail(reciepents, "Test Send Email", "Body Goes Here")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

```