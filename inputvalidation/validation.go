package inputvalidation

import (
	"net/url"
	"regexp"
)

type Student struct {
	StudentID string
	Name      string
	Email     string
}

func (a Student) IsValid() url.Values {
	errs := url.Values{}

	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// check if the studentid empty

	if a.StudentID == "" {

		errs.Add("studentid", "The StudentID is not empty!")
	}

	if a.Name == "" {

		errs.Add("name", "The name is required!")
	}

	// check the name field is between 2 to 50 characters
	if len(a.Name) < 2 || len(a.Name) > 50 {

		errs.Add("name", "The name field must be between 2-50 chars!")
	}

	if a.Email == "" {

		errs.Add("email", "The email field is required!")
	}

	if !regexpEmail.Match([]byte(a.Email)) {

		errs.Add("email", "The email field should be a valid email address!")
	}
	return errs
}
