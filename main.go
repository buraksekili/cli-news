package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/mmcdole/gofeed"
	"os"
	"strings"
)

type Content struct {
	URL      string
	Headline string
}

func main() {
	validFlag, cont := parseFlag()
	if !validFlag || !cont {
		os.Exit(1)
	}

	fp := gofeed.NewParser()
	var URL string = "https://news.ycombinator.com/rss"
	feed, _ := fp.ParseURL(URL)

	m := make(map[string]string)
	itemsArr := feed.Items
	contentArr := make([]Content, len(itemsArr))

	for i := 0; i < len(itemsArr); i++ {
		if itemsArr[i] != nil {
			m[strings.Trim(itemsArr[i].Title, " ")] = itemsArr[i].Link
			contentArr[i] = Content{Headline: itemsArr[i].Title, URL: itemsArr[i].Link}
		}
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "> {{ .Headline | green | bold }} ",
		Inactive: " {{ .Headline | white }} ",
		Details: `
		--------- News ----------
		{{ "Headline:" | faint }}	{{ .Headline }}
		{{ "Link:" | faint }}	{{ .URL}}`,
		Help: " ",
		// delete default help text
		// as indicated in https://github.com/manifoldco/promptui/blob/981a3cab68f6f3481bf42c6a98521af7fbd14fae/select.go#L472
	}

	prompt := promptui.Select{
		Label:     "Select News",
		Items:     contentArr,
		Templates: templates,
		Size:      10,
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
		PrintHelp(os.Stdout)
		return true, false
	}

	err := "Unrecognized flag " + op
	fmt.Printf("%s\n", color.RedString(err))
	return false, false
}
