package cmd

import (
	"github.com/knet-network/pw-exchange/pkg/credentials"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

const ArgDelimiter = "="

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generatee a link to exchange a password",
	Run:   generate,
}

// sth like generate KEY=VALUE KEY=VALUE should be used
// todo move in package credentials
func generate(cli *cobra.Command, args []string) {
	if len(args) <= 0 {
		//sry there is nothing to do
		_ = cli.Usage()
	}

	values := make([]credentials.KeyValue, len(args))
	for _, arg := range args {
		//should consist of key=value pairs
		splits := strings.Split(arg, ArgDelimiter)
		if len(splits) != 2 {
			log.Fatalf("argument %s is invalid", arg)
		}
		keyValue := credentials.KeyValue{Key: splits[0], Value: splits[1]}
		values = append(values, keyValue)
	}

	//now create the link
	credentials.NewInfoEntry(values)
}
