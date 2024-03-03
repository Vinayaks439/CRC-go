/*
Copyright © 2024 Vinayak <vinayaks23@iitk.ac.in>
*/
package cmd

import (
	"CNassignment/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	data string
	key  string
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode [flags] data [...data] key [...key]",
	Short: "This command encodes the Data provided with CRC tech",
	Long: `This Program applies CRC encoding to a provided data and a key by doing modulo2 division with XOR 
		Operation. The generated remainder is added back to the data and sent as a Frame. 
		To use this program provide data and key as commandline args`,
	Run: func(cmd *cobra.Command, args []string) {
		frame := pkg.CRC(data, key)
		fmt.Printf("The Encoded DataFrame after applying CRC is %s\n", frame)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	flags := encodeCmd.PersistentFlags()
	flags.StringVarP(&data, "data", "d", "", "Please Enter the DATA to be encoded using -d or --data")
	flags.StringVarP(&key, "key", "k", "", "Please Enter the Key for encoding using -k")
}
