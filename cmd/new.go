package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcnksm/go-input"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		ui := &input.UI{Writer: os.Stdout, Reader: os.Stdin}

		shopDomain, err := ui.Ask("Shop Domain", &input.Options{
			Required:  true,
			Loop:      true,
			HideOrder: true,
		})
		if err != nil {
			return err
		}

		password, err := ui.Ask("Password", &input.Options{
			Required:  true,
			Loop:      true,
			Mask:      true,
			HideOrder: true,
		})
		if err != nil {
			return err
		}

		fmt.Println("output:", shopDomain, password)
		return nil
	},
}
