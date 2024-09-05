package ascii

import (
	"log"
	"os"
	"strings"
)

var Banner string

func BannerManagement() string {
	align = AlignmentManagement()
	args := os.Args[1:]
	// Usage instructions if no arguments or more than 2 arguments are provided
	if len(args) == 0 || len(args) > 3 {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}
	// the default styling will be standard.txt unless the user chooses otherwise
	Banner = "standard.txt"
	if len(args) == 1 {
		args = append(args, Banner)
	}
	if len(args) == 3 {
		Banner = args[2]
	}
	if len(args) == 2 && (args[0] == "--align=left" || args[0] == "--align=right" || args[0] == "--align=center" || args[0] == "--align=justify") {
		input = args[1]
		align = args[0]
		args = append(args, Banner)
	}

	// we add a .txt extension to the banner if it doesn't already exist
	// this way weather the user adds it or not it won't make a difference
	if !strings.HasSuffix(Banner, ".txt") {
		Banner += ".txt"
	}
	// we inform the user of the available styles if he doesn't choose a proper one
	if Banner != "standard.txt" && Banner != "thinkertoy.txt" && Banner != "shadow.txt" {
		log.Fatal("Usage: this style is unavailable \nPlease choose one of the available styles \n1 : standard \n2 : thinkertoy \n3 : shadow")
	}
	return Banner
}
