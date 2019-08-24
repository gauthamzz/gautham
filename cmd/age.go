package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// ageCmd represents the age command
var ageCmd = &cobra.Command{
	Use:   "age",
	Short: "Return age",
	Long:  `Age of Gautham`,
	Run: func(cmd *cobra.Command, args []string) {
		format := "Jan 2, 2006" // just a format. not taken into calculation
		DOB := "Sep 10, 1995"
		start, _ := time.Parse(format, DOB)
		end := time.Since(start)
		years := end / time.Hour / 24 / 365

		fmt.Printf("Gautham is %d years old 💛 \n", years)
	},
}

func init() {
	rootCmd.AddCommand(ageCmd)

}
