package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// whyCmd represents the why command
var whyCmd = &cobra.Command{
	Use:   "why",
	Short: "Why did i make this?",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The reason anyone would do this if they could, which they can't would be because they could, which they can't. \n\n https://www.youtube.com/watch?v=lbKKfFum_iw")
	},
}

func init() {
	rootCmd.AddCommand(whyCmd)

}
