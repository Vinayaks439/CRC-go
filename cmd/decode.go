/*
Copyright © 2024 Vinayak S <vinayaks23@iitk.ac.in>
*/
package cmd

import (
	"CNassignment/pkg"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	frame string
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checksum := pkg.CheckSum(frame, key)
		fmt.Printf("The Checksum after decoding the frame with the key is %s\n", checksum)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	flags := decodeCmd.PersistentFlags()
	flags.StringVarP(&frame, "frame", "f", "", "Please Enter the Frame to be decoded using -f or --frame")
	flags.StringVarP(&key, "key", "k", "", "Please Enter the CRC Key for decoding using -k")
}
