package email

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/labstack/echo/v4"
	"sittellalab.com.au/templates"
)

var EmailTemplate = "From: <%s>\nTo: <%s>\nSubject: %s\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s"

func SendMail(ctx echo.Context) error {
	// parse email form
	enquiry, err := ParseContactForm(ctx)
	if err != nil {
		return err
	}

	// authenticate email login
	auth := smtp.PlainAuth("", os.Getenv("ICLOUD_USER"), os.Getenv("ICLOUD_PASS"), os.Getenv("SMTP"))

	// send the acknowledgement email
	var resBuf bytes.Buffer
	if err := templates.EnquiryResponse(enquiry.Templ()).Render(ctx.Request().Context(), &resBuf); err != nil {
		log.Println(err)
		return fmt.Errorf(Error)
	}
	response := fmt.Sprintf(EmailTemplate, os.Getenv("EMAIL"), enquiry.Email, "Sittella Lab Enquiry", resBuf.String())

	if err := smtp.SendMail(
		os.Getenv("SMTP")+":"+os.Getenv("SMTP_PORT"),
		auth,
		os.Getenv("EMAIL"),
		[]string{enquiry.Email},
		[]byte(response),
	); err != nil {
		log.Println(err)
		return fmt.Errorf(Error)
	}

	// send the enquiry to the hello@sittellalab.com.au email
	var enqBuf bytes.Buffer
	if err := templates.EnquiryEmail(enquiry.Templ()).Render(ctx.Request().Context(), &enqBuf); err != nil {
		log.Println(err)
		return fmt.Errorf(Error)
	}
	forward := fmt.Sprintf(EmailTemplate, os.Getenv("EMAIL"), os.Getenv("EMAIL"), "Website Enquiry", enqBuf.String())

	if err := smtp.SendMail(
		os.Getenv("SMTP")+":"+os.Getenv("SMTP_PORT"),
		auth,
		os.Getenv("EMAIL"),
		[]string{os.Getenv("EMAIL")},
		[]byte(forward),
	); err != nil {
		log.Println(err)
		return fmt.Errorf(Error)
	}

	return nil
}
