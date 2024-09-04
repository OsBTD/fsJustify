package ascii

import (
	"log"
	"os"
	"strings"
)

var (
	input      string
	inputsplit []string
)

func InputManagement() (string, []string) {
	align = AlignmentManagement()
	Banner = BannerManagement()
	// here we'll read the text the user inputs
	args := os.Args[1:]
	if len(args) == 1 {
		input = args[0]
	} else if len(args) > 1 {
		input = args[1]
	}
	// we split it by literal newline
	inputsplit = strings.Split(input, "\\n")
	// we check if the input is valid and only contains printable ascii characters
	for _, line := range inputsplit {
		for _, c := range line {
			if c < 32 || c > 126 {
				log.Fatal("Error : input should only contain printable ascii characters")
			}
		}
	}
	// we return the raw input and the one split by literal newlines
	return input, inputsplit
}
