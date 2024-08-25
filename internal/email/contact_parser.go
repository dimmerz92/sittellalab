package email

import (
	"fmt"
	"log"
	"net/mail"
	"os"

	"github.com/labstack/echo/v4"
	"sittellalab.com.au/templates"
)

type Enquiry struct {
	FirstName   string
	LastName    string
	Email       string
	EnquiryType string
	Message     string
}

func (e *Enquiry) Templ() templates.Enquiry {
	return templates.Enquiry{
		Url:         os.Getenv("URL"),
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		Email:       e.Email,
		EnquiryType: e.EnquiryType,
		Message:     e.Message,
	}
}

var Error = "It looks like we're having issues, please try again later."

func ParseContactForm(ctx echo.Context) (*Enquiry, error) {
	form, err := ctx.FormParams()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(Error)
	}

	firstName, ok := form["firstname"]
	if !ok || len(firstName[0]) < 1 {
		return nil, fmt.Errorf("First name required")
	}

	lastName, ok := form["lastname"]
	if !ok || len(lastName[0]) < 1 {
		return nil, fmt.Errorf("Last name required")
	}

	email, ok := form["email"]
	if !ok || len(email[0]) < 1 {
		return nil, fmt.Errorf("Email required")
	} else if _, err := mail.ParseAddress(email[0]); err != nil {
		return nil, fmt.Errorf("Email does not appear to be correct")
	}

	enquiryType, ok := form["enquiry-type"]
	if !ok || len(enquiryType[0]) < 1 {
		return nil, fmt.Errorf("Enquiry type required")
	}

	message, ok := form["message"]
	if !ok || len(message[0]) < 1 {
		return nil, fmt.Errorf("Message required")
	}

	return &Enquiry{
		FirstName:   firstName[0],
		LastName:    lastName[0],
		Email:       email[0],
		EnquiryType: enquiryType[0],
		Message:     message[0],
	}, nil
}
