package main

import (
	"log"
	"os"

	"github.com/herux/hxmail"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal("too few arguments. Run ./cmd username@account.com smtp_password to_email smtp_host")
	}

	username := os.Args[1]
	password := os.Args[2]
	to := os.Args[3]
	host := os.Args[4]
	email := hxmail.NewHxMail(
		host,
		587,
		username,
		password)
	email.To(to)
	email.Subject("test hxmail")
	email.SendMail()
}
