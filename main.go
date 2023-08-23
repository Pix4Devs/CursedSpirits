package main

import (
	"Pix4Devs/CursedSpirits/cmd"
	"Pix4Devs/CursedSpirits/fancy"
	"fmt"
	"log"
)

func main() {
	logo := fancy.ConcatLogo()
	logo.Colorize()

	fmt.Println(*logo + "\n\n")

	cmd.Init()
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
