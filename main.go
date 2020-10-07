package main

import (
	"fmt"
	"github.com/buraksekili/cli-news/scraper"
	"github.com/manifoldco/promptui"
)

func main() {
	allContents, _ := scraper.GetHackerNews()
	m := make(map[string]string)
	for _, content := range allContents {
		m[content.Headline] = content.URL
		fmt.Printf("Headline: %s\nLink: %s\n\n", content.Headline, content.URL)
	}

	keys := make([]string, len(m))
	i := 0
	for key, _ := range m {
		keys[i] = key
		i++
	}

	prompt := promptui.Select{
		Label: "Select News ('q' to quit)",
		Items: keys,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose: %q\nLink: %s\n", result, m[result])
}
