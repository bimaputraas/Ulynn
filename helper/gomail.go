package helper

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(email string,userId int,codestring string) error{
	m := gomail.NewMessage()
	m.SetHeader("From", "ddummymail65@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Registration Successful")
	body := fmt.Sprintf("Congratulations! Your registration on the Video Game Rental platform was successful. Thank you for choosing Video Game Rental and happy gaming!\n\nYour verification link : http://localhost:8080/users/statusverification/%d/%s",userId,codestring)
	m.SetBody("text/html", body)
	
	password := os.Getenv("PASSWORDSMTP")
	d := gomail.NewDialer("smtp.elasticemail.com", 2525, "ddummymail65@gmail.com",password)
	
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}