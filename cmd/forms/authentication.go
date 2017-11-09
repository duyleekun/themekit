package forms

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey"
)

// Auth holds the values returned from the auth form
type Auth struct {
	Domain   string
	Login    string
	Password string
}

var authForm = []*survey.Question{
	{
		Name: "domain",
		Prompt: &survey.Input{
			Message: "Shopify Domain",
			Help:    "Your shop domain ending with .myshopify.com",
		},
		Validate: myShopifyDomain,
	},
	{
		Name: "login",
		Prompt: &survey.Input{
			Message: "Login",
			Help:    "Email you use to login to shopify",
		},
		Validate: isEmail,
	},
	{
		Name: "password",
		Prompt: &survey.Password{
			Message: "Password",
			Help:    "Your shopify login password",
		},
		Validate: survey.Required,
	},
}

// RequestAuthDetails returns authentication details for logging into shopify
func RequestAuthDetails() (*Auth, error) {
	auth := &Auth{}
	err := survey.Ask(authForm, auth)
	return auth, err
}

func myShopifyDomain(val interface{}) error {
	domain := val.(string)

	if len(domain) == 0 {
		return fmt.Errorf("missing store domain")
	} else if !strings.HasSuffix(domain, "myshopify.com") &&
		!strings.HasSuffix(domain, "myshopify.io") &&
		!strings.HasPrefix(domain, "http://127.0.0.1:") { // for testing
		return fmt.Errorf("invalid store domain must end in '.myshopify.com'")
	}

	return nil
}

func isEmail(val interface{}) error {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	email := val.(string)

	if !emailRegexp.MatchString(email) {
		return fmt.Errorf("Invalid email")
	}

	return nil
}
