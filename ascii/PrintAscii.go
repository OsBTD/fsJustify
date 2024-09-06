package ascii

import (
	"fmt"
)

func PrintArt(string, []string, map[rune]([]string)) {
	TerminalLength, AsciiLength, NoSpaces, countspace = CalculateLength()
	align = AlignmentManagement()
	// we start by checking if the user input only contains literal newlines
	// if so we print newlines accordingly
	for idx, line := range inputsplit {
		// also if there's empty strings resulting from the spliting we print a newline
		if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
			fmt.Println()
			continue
		} else if len(line) == 0 && !Checknewline(inputsplit) {
			fmt.Println()
		} else if len(line) != 0 && !Checknewline(inputsplit) {
			if align == "--align=left" {
				for i := 0; i < 8; i++ {
					for j := 0; j < len(line); j++ {
						inputrune := rune(line[j])

						fmt.Print(Replace[inputrune][i])

					}
					// after each line is printed we print a newline
					fmt.Println()
				}
			} else if align == "--align=right" {
				for i := 0; i < 8; i++ {
					for l := 0; l < (TerminalLength - AsciiLength); l++ {
						fmt.Print(" ")
					}

					for j := 0; j < len(line); j++ {
						inputrune := rune(line[j])

						fmt.Print(Replace[inputrune][i])

					}
					// after each line is printed we print a newline
					fmt.Println()
				}
			} else if align == "--align=center" {
				for i := 0; i < 8; i++ {
					for l := 0; l < (TerminalLength/2 - AsciiLength/2); l++ {
						fmt.Print(" ")
					}

					for j := 0; j < len(line); j++ {
						inputrune := rune(line[j])

						fmt.Print(Replace[inputrune][i])

					}
					// after each line is printed we print a newline
					fmt.Println()
				}
			} else if align == "--align=justify" {
				for i := 0; i < 8; i++ {

					for j := 0; j < len(line); j++ {
						if line[j] != 32 {
							inputrune := rune(line[j])

							fmt.Print(Replace[inputrune][i])

						} else {
							for l := 0; l < (TerminalLength-NoSpaces)/countspace; l++ {
								fmt.Print(" ")
							}
						}
					}

					// after each line is printed we print a newline
					fmt.Println()
				}
			}
		}
	}
}
