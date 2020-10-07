package scraper

import (
	"fmt"
	"github.com/fatih/color"
	"io"
)

func PrintHelp(w io.Writer) {
	name := "------ cli-news ------:\n "
	help := "cli-news is a Hacker news scraper to display popular headlines of Hacker news without using external API to fetch popular headlines.\n\n" +
		"-h, --help boolean\n\tPrints this message.\n" +
		"<C-c> to quit program.\n"
	fmt.Fprintf(w, "%s%s\n", color.GreenString(name), help)
}
