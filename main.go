package main

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func date() string {
	dt := time.Now()
	//Format dd-mm
	return (dt.Format("02-01"))
}

func sendMail(to []string, from, password, subject, body string) {
	smtpHost := "smtp.mail.yahoo.com"
	smtpPort := "587"

	// Message.
	message := []byte(subject + "\n" + "This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func wishBirthday() {
	f, err := excelize.OpenFile(`C:\Users\Lenovo\Documents\MS Office\Excel files\birthdays.xlsx`)
	if err != nil {
		panic(err)
	}

	cols, err := f.GetCols("Sheet1")
	if err != nil {
		panic(err)
	}

	names := cols[0]
	emails := cols[2]
	birthdays := cols[1]


	for i := range names {
		if i == 0 {
			continue
		}

		name := names[i]
		email := emails[i]
		birthday := birthdays[i]

		if birthday == date() {
			subject := "Happy Birthday " + name
			body := fmt.Sprintf("Hey %v, \nHappy birthday to you! Wish you an octotastic life ahead!\nYours lovely,\nShravan", name)
			sendMail([]string{email}, "delunator.one@yahoo.com", "ytxyqkendtnngqib", subject, body)
			fmt.Println("Wished happy birthday to", name)
		
		} else {
			fmt.Println("No birtday of", name)
		}
	}
}


func main() {
	for {
		wishBirthday()
		fmt.Println("Sleeping for a day...")
		time.Sleep(time.Second * 86400)
	}
}