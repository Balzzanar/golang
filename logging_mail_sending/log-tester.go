package main

import (
    "github.com/apsdehal/go-logger"
    "github.com/go-gomail/gomail"
    "os"
)

var log *logger.Logger

func main () {
    var errr error
    log, errr = logger.New("woop", 1, os.Stdout)
    if errr != nil {
        panic(errr) // Check for error
    }

    log.Info("Sending mail...")
    send_mail()
}



func send_mail(){
	m := gomail.NewMessage()
	m.SetHeader("From", "toby.tooth@mail.com")
	m.SetHeader("To", "toby.tooth@mail.com", "toby.tooth+1@mail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Toby</b> and <i>Toby +1</i>!")

	d := gomail.NewDialer("smtp.mail.com", 465, "toby.tooth@mail.com", "****")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Error("Mail send faild!")
	    panic(err)
	}
	log.Info("Mail sent!")
}