package fancy

import (
	"math/rand"
	"strings"
)

type LogoStr string

func ConcatLogo() *LogoStr {
	var logo []string
	var lg LogoStr

	logo = append(logo, "    ╔═╗╦ ╦╦═╗╔═╗╔═╗╔╦╗  ╔═╗╔═╗╦╦═╗╦╔╦╗╔═╗")
	logo = append(logo, "    ║  ║ ║╠╦╝╚═╗║╣  ║║  ╚═╗╠═╝║╠╦╝║ ║ ╚═╗")
	logo = append(logo, "    ╚═╝╚═╝╩╚═╚═╝╚═╝═╩╝  ╚═╝╩  ╩╩╚═╩ ╩ ╚═╝")

	lg = LogoStr(strings.Join(logo,"\n"))
	return &lg
}

func (logo *LogoStr) Colorize() {
	chain := strings.Split(string(*logo), "\n")

	var out []string
	for _, v := range chain {
		var buff []string
		for _, v_ := range v {
			color := colors[rand.Intn(len(colors))]
			buff = append(buff, color+string(+v_)+"\x1b[0m")
		}

		out = append(out, strings.Join(buff, ""))
	}

	*logo = LogoStr(strings.Join(out, "\n"))
	{
		*logo = *logo + "\n" + "\x1b[1m\x1b[38;5;197mSearching for someone to blame is such a pain.\x1b[0m"
		*logo = *logo + "\n" + "               \x1b[1m\x1b[38;5;63m@\x1b[38;5;160mz3ntl3 \x1b[38;5;63m& \x1b[38;5;160mmidas\x1b[0m "
	}
}

var (
	colors = []string{
		"\x1b[38;5;197m",
		"\x1b[38;5;196m",
		"\x1b[38;5;160m",
		"\x1b[38;5;161m",
		"\x1b[38;5;125m",
		"\x1b[38;5;124m",
	}
)