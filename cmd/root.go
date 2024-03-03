/*
Copyright © 2024 Vinayak S <vinayaks23@iitk.ac.in>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "./crc encode [flags]",
	Short: "This Program generates a data frame after apply CRC encoding",
	Long: `This Program applies CRC encoding to a provided data and a key by doing modulo2 division with XOR 
		Operation. The generated remainder is added back to the data and sent as a Frame. 
		To use this program provide data and key as commandline args`,
	//Run: func(cmd *cobra.Command, args []string) {
	//
	//},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
