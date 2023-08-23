package main

import (
	"Pix4Devs/CursedSpirits/cmd"
	"Pix4Devs/CursedSpirits/fancy"
	"flag"
	"fmt"
	"log"
)

var (
	TARGET      = flag.String("url", "", "Target URL. Examples: https://github.com or http://google.com")
	CONCURRENCY = flag.Int("concurrency", 2000, "Defines concurrency across requests")
	DURATION    = flag.Int("duration", 300, "Flood duration in seconds")
	PROXY_TYPE = flag.String("protocol", "socks4", "Proxy protocol, can be one of: socks4 or socks5")
	VALID_PROTOCOLS = []string{"socks4","socks5"}
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
