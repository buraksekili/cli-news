package main

import (
	"fmt"
	"github.com/buraksekili/cli-news/scraper"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func main() {
	validFlag, cont := parseFlag()
	if !validFlag || !cont {
		os.Exit(1)
	}

	allContents, _ := scraper.GetHackerNews()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Headline | green | bold }} ",
		Inactive: " {{ .Headline | white }} ",
		Selected: "	{{ .Headline | green | cyan }}",
		Details: `
	--------- News ----------
	{{ "Headline:" | faint }}	{{ .Headline }}
	{{ "URL:" | faint }}	{{ .URL}}`,
	}

	prompt := promptui.Select{
		Label:     "Select News",
		Items:     allContents,
		Size:      10,
		Templates: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	i := strings.Index(result, " ")
	fmt.Printf("\nYou choose: %q\nLink: %s\n", result[i+1:len(result)-1], result[1:i])
}

// parseFlag() parses input flags and returns
// 2 boolean. First one validates flag and second one
// indicates whether program continue to run or not.
func parseFlag() (bool, bool) {
	var flags []string = os.Args[1:]
	if len(flags) > 1 {
		return false, false
	}
	if len(flags) == 0 {
		return true, true
	}

	var op string = flags[0]
	if op == "-h" || op == "--help" {
		scraper.PrintHelp(os.Stdout)
		return true, false
	}

	err := "Unrecognized flag " + op
	fmt.Printf("%s\n", color.RedString(err))
	return false, false
}
