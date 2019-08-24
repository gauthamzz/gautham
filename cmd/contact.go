package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Contact struct {
	label string
	url   string
}

type Contacts []Contact

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Links to website and social media",
	Long:  `Website and social`,
	Run: func(cmd *cobra.Command, args []string) {
		contacts := getContacts()
		for _, contact := range contacts {
			fmt.Print(
				contact.label,
				": ",
				contact.url,
				"\n",
			)
		}
	},
}

func getContacts() Contacts {
	contacts := Contacts{
		{
			"Website",
			"https://www.gauthamzz.com",
		},
		{
			"GitHub",
			"https://github.com/gauthamzz",
		},
		{
			"Twitter",
			"https://twitter.com/gauthamzzz",
		},
		{
			"Medium",
			"https://medium.com/@gauthamsanthosh",
		},
		{
			"Facebook",
			"https://www.facebook.com/gauthamzz",
		},
		{
			"Instagram",
			"https://www.instagram.com/gauthamzz/",
		},
	}
	return contacts
}
func init() {
	rootCmd.AddCommand(contactCmd)
}
