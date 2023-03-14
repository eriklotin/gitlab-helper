package main

import (
	"fmt"
	"github.com/eriklotin/gitlab-helper/client"
	"github.com/eriklotin/gitlab-helper/config"
)

func main() {
	conf := config.Init()
	gitlabClient := client.GetClient(conf.Token)

	listOfMR := gitlabClient.GetMyOpenedMRs()

	fmt.Println("Your open merge requests:")
	fmt.Println()
	for _, m := range listOfMR {
		fmt.Println("* ", m.Title)
		fmt.Println("    ", m.WebUrl)
		fmt.Println()
	}
}
