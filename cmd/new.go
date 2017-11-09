package cmd

import (
	"fmt"

	"github.com/Shopify/themekit/cmd/forms"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		auth, err := forms.RequestAuthDetails()
		if err != nil {
			return err
		}

		fmt.Printf(
			"Check Auth\nDomain:%s\nLogin:%s\nPass:%s\n",
			auth.Domain,
			auth.Login,
			auth.Password,
		)

		action, err := forms.RequestAction(
			[]string{"timber1", "brooklyn"},
			[]string{"Timber", "Brooklyn"},
		)
		if err != nil {
			return err
		}

		fmt.Printf(
			"Theme Action\ntheme:%s\nduplicate:%v\nname:%s\n",
			action.Theme,
			action.Duplicate,
			action.Name,
		)

		return nil
	},
}
