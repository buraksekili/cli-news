package main

import (
	"fmt"
	"github.com/buraksekili/cli-news/scraper"
	"github.com/manifoldco/promptui"
	"strings"
)

func main() {
	allContents, _ := scraper.GetHackerNews()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   " {{ .Headline | green | bold }} ",
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
