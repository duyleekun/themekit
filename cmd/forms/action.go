package forms

import (
	"github.com/AlecAivazis/survey"
)

// Action definitions
const (
	createNewTheme    = "Create a new theme"
	workOnExisting    = "Work on an existing theme on Shopify"
	duplicateExisting = "Duplicate existing theme on Shopify"
)

// ThemeAction describes which action will be taken after authentication
type ThemeAction struct {
	Theme     string
	Duplicate bool
	Name      string
}

var actionSelect = &survey.Select{
	Message: "What would you like to do",
	Options: []string{createNewTheme, workOnExisting, duplicateExisting},
}

// RequestAction prompt the user for which action they would like to take while
// running the new command
func RequestAction(existing, templates []string) (*ThemeAction, error) {
	var action string
	err := survey.AskOne(actionSelect, &action, nil)
	if err != nil {
		return nil, err
	}

	actionForm := []*survey.Question{}
	switch action {
	case createNewTheme:
	case workOnExisting:
		actionForm = []*survey.Question{
			{
				Name: "theme",
				Prompt: &survey.Select{
					Message: "Select a theme to work on",
					Options: existing,
				},
			},
		}
	case duplicateExisting:
		actionForm = []*survey.Question{
			{
				Name: "theme",
				Prompt: &survey.Select{
					Message: "Select a theme to work on",
					Options: existing,
				},
			},
			{
				Name: "name",
				Prompt: &survey.Input{
					Message: "Name for this theme (optional)",
					Help:    "This will create change the name of this theme in the Shopify Admin",
				},
			},
		}
	}

	themeAction := &ThemeAction{}
	err = survey.Ask(actionForm, themeAction)
	themeAction.Duplicate = duplicateExisting == action
	return themeAction, err
}
