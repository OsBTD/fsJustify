package ascii

import (
	"log"
	"os"
)

var align string

func AlignmentManagement() string {
	args := os.Args[1:]
	// Usage instructions if no arguments or more than 2 arguments are provided
	if len(args) == 0 || len(args) > 3 {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}
	align = "--align=left"
	if len(args) == 1 {
		args = append(args, align)
	}

	if len(args) == 3 {
		align = args[0]
	}

	if len(args) == 2 && (args[1] == "standard.txt" || args[1] == "shadow.txt" || args[1] == "thinkertoy.txt" || args[1] == "standard" || args[1] == "shadow" || args[1] == "thinkertoy") {
		input = args[0]
		Banner = args[1]
		args = append(args, align)
	}
	// the default alignment will be left unless the user chooses otherwise
	if align == "left" || align == "right" || align == "justify" || align == "center" {
		align = "--align=" + align
	}
	if align != "--align=left" && align != "--align=right" && align != "--align=justify" && align != "--align=center" {
		log.Fatal("Usage: this alignment is unavailable \nPlease choose one of the available alignments \n1 : left \n2 : right \n3 : justify")
	}

	if len(args) == 1 {
		args = append(args, align, Banner)
	} else if len(args) == 2 && args[0] == align {
		args = append(args, Banner)
	}
	// we formate the alignment in the correct way if the user doesn't do so
	// this way weather the user does the formatting correctly or not it won't make a difference
	// we inform the user of the available styles if he doesn't choose a proper one
	return align
}
