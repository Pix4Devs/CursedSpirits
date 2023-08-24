package cmd

import (
	"Pix4Devs/CursedSpirits/filesystem"
	"Pix4Devs/CursedSpirits/globals"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/MidasVanVeen/proxy-checker/pkg/proxychecker"
	"github.com/MidasVanVeen/proxy-checker/pkg/proxyutils"
	"github.com/Pix4Devs/pix4lib"
	"github.com/Pix4Devs/pix4lib/proxyscraper"
	"github.com/corpix/uarand"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "./CursedSpirits",
}

var SubCommands = []*cobra.Command{
	{
		Use: "start",
		Short: "Start flood",
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
			defer os.Exit(0)
			fmt.Printf("[VERSION] %s", globals.VERSION)
		},
	},
	{
		Use: "check",
		Short: "Check utilities",
	},
}


func Init(){
	RootCmd.AddCommand(SubCommands...)
	/*
		Any command and sub commands other than 'start' need to exit the process after the command has been finished!
		Keep that in mind...
	*/
	{
		// FLAGS for command 'start'
		SubCommands[0].Flags().StringVar(&globals.TARGET,"url","","Sets target URL, requires http:// or https:// scheme included!")
		SubCommands[0].MarkFlagRequired("url")

		{
			SubCommands[0].Flags().IntVar(&globals.CONCURRENCY,"concurrency",2000,"Set concurrency accross requests")
			SubCommands[0].Flags().IntVar(&globals.DURATION,"duration",300,"Sets flood duration in seconds")
			SubCommands[0].Flags().StringVar(&globals.PROXY_TYPE,"protocol","socks4","Sets proxy type, can be one of socks4 or socks5")
		}
		
		// FLAGS for command 'scrape'
		SubCommands[1].Flags().Int("timeout", 1000, "Defines timeout in seconds for proxy scraping")

		// Add subcommand for 'scrape' parent command 'scrape'
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

		// Add subcommand 'proxy' to parent command 'check'
		proxy := cobra.Command{
			Use: "proxy",
			Short: "Proxy checker supports only socks4/socks5",
			Run: func(cmd *cobra.Command, args []string) {
				defer os.Exit(0) // so that it does not bypass app break after this command has been performed,
				checkProxies(cmd)
			},
		}
		(&proxy).Flags().Int("timeout", 5, "Defines timeout in seconds for proxy checking")
		(&proxy).Flags().Int("retries", 3, "Defines how many times the checker should retry to check a proxy")
		SubCommands[3].AddCommand(&proxy)
		// Below here add required flags if needed
		// TODO
	}
}

func checkProxies(cmd *cobra.Command) {
	tmOut, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		tmOut = 5
	}
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
	retries, err := cmd.Flags().GetInt("retries")
	if err != nil {
		retries = 3
	}
	checker, err := proxychecker.NewChecker(proxyutils.SOCKS5, time.Second * time.Duration(tmOut), retries);
	if err != nil {
		log.Fatal(err)
	}
	uas := make([]string, 0)
	for i := 0; i < len(globals.PROXIES); i++ {
		uas = append(uas, uarand.GetRandom())
	}
	i:=0
	cleanProxies := checker.CleanList(globals.PROXIES, &uas, &globals.REFS, func(proxy string) {
		i++
		fmt.Printf("\r[PROXY CHECKER] %s\t[%d/%d]", proxy, i, len(globals.PROXIES));
	})
	fmt.Println()
	// ask user if he wants to save the checked proxies
	var save string
	fmt.Print("[PROXY CHECKER] Enter file path for valid proxies. Leave blank to overwrite proxies.txt: ")
	fmt.Scanln(&save)
	// TODO: validate file path
	saveFile := filepath.Join(base,"context","proxies.txt")
	if save != "" {
		saveFile = filepath.Join(base, save)
	}
	file, err := os.OpenFile(saveFile, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(strings.Join(cleanProxies, "\n"))
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
