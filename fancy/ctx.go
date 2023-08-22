package fancy

import (
	"fmt"
	"os"
	"strings"
)

func PrintCtx(ctx string) {
	var table []string

	table = append(table, "\n\n\x1b[38;5;160m╔═══════\x1b[38;5;161m═══\x1b[38;5;197m➤\x1b[0m")
	table = append(table, fmt.Sprintf("\x1b[38;5;124m║ \x1b[0m \x1b[38;5;197m%s", ctx))
	table = append(table, "\x1b[38;5;125m╚\x1b[38;5;124m══\x1b[38;5;161m═\x1b[38;5;197m➤\x1b[0m")

	fmt.Fprint(os.Stdout, strings.Join(table, "\n"))
}