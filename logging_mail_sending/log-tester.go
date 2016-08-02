package main

import (
    "github.com/apsdehal/go-logger"
    "github.com/go-gomail/gomail"
    "os"
    "io/ioutil"
    "flag"
    "fmt"
    "errors"
)

var log *logger.Logger
var filename *string

func main () {
    var err error
    log, err = logger.New("woop", 1, os.Stdout)
    if err != nil {
        panic(err) 
    }

    if err = init_args(); err != nil {
    	panic(fmt.Sprintf("Bad argument input! (%s)\nuse '--help' to get a list of arguments", err))
    }
    log.Info("Reading file...")
    content := read_file(*filename)
    log.Info("Sending mail...")
  	send_mail(content)
  	log.Info("Mail sent!")
}

func init_args() error {
	var result error 
	filename = flag.String("f", "", "Filename of the result file, that shall be mailed")
	flag.Parse()
	if _, err := os.Stat(*filename); os.IsNotExist(err) {
 		log.Error(fmt.Sprintf("The file '%s' does not exist!", *filename))
		result = errors.New("missing filelocation")
	}
	return result
}

func read_file(filename string) []byte {
	content,_ := ioutil.ReadFile(filename)
	return content
}


func send_mail(content []byte) {
	m := gomail.NewMessage()
	m.SetHeader("From", "toby.tooth@mail.com")
	m.SetHeader("To", "spann.johan@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", 
		"<h1>Result</h1>" + 
		"<p>"+ string(content) +"</p>")
	d := gomail.NewDialer("smtp.mail.com", 465, "toby.tooth@mail.com", "****")
	if err := d.DialAndSend(m); err != nil {
		log.Error("Mail send faild!")
	    panic(fmt.Sprintf("Mailing failed! (%s)", err))
	}
}