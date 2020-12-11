package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

var (
	FROMADDR string // user@example.com
	PASSWORD string // password
	SERVER string // smtp.example.com
	PORT string // 25
	RECIPIENTS string // comma delimited "user@example.com,user2@another.com"
	MSGFILE string // path to file containing RFC822 (tools.ietf.org/html/rfc822) message
)

func init() {
	flag.StringVar(&FROMADDR, "from", "me@example.com", "sender account")
	flag.StringVar(&PASSWORD, "pwd", "mypassword", "password for account")
	flag.StringVar(&SERVER, "server", "mail.example.com", "mail server")
	flag.StringVar(&PORT, "port", "25", "mail server smtp port")
	flag.StringVar(&RECIPIENTS, "to", "user1@this.com,user2@that.com", "comma-delimited recipients")
	flag.StringVar(&MSGFILE, "msg", "message.txt", "RFC 822-style message")

	flag.Parse()
}

func main() {
	// Set up logger
	logFile, _ := os.Create("mta.log")
	defer logFile.Close()
	logger := log.New(logFile, "jbmta", 0)

	// Open message file for reading
	f, err := os.Open(MSGFILE)
	defer f.Close()
	if err != nil {
		logger.Fatal(err)
	}

	// Authenticate with the server
	auth := smtp.PlainAuth("", FROMADDR, PASSWORD, SERVER)
	to := strings.Split(RECIPIENTS, ",")
	msg, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Fatal(err)
	}

	// Send mail
	if err := smtp.SendMail(SERVER+":"+PORT, auth, FROMADDR, to, msg); err != nil {
		logger.Fatal(err)
	}
}
