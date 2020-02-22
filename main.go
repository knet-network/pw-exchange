package main

import (
	"fmt"
	"github.com/knet-network/pw-exchange/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("An error occurred!", err)
	}
}
