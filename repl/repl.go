package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rokob/kudu/parser"
)

// Run - run the repl
func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	showTagline()
	showInputLine()
	var currentInput string
	for scanner.Scan() {
		lastInput := scanner.Text()
		currentInput += lastInput
		ok, isBreak, output := handleLineOfText(currentInput)
		if !ok && !isBreak {
			okLast, isBreakLast, _ := handleLineOfText(lastInput)
			if !okLast && isBreakLast {
				currentInput = ""
				showInputLine()
			} else {
				currentInput += "\n"
				showContinuationLine()
			}
		} else {
			if ok {
				showOutput(output)
				currentInput = ""
				showInputLine()
			} else if isBreak {
				currentInput = ""
				showInputLine()
			} else {
				currentInput += "\n"
				showContinuationLine()
			}
		}
	}
	showExit()
}

func showTagline() {
	fmt.Println("Kudu v0.1 - The Kuduest Language")
	fmt.Println("If you screw up, type $ to end this expression")
	fmt.Println("^D to quit")
}

func showExit() {
	fmt.Println("\n>>> Kudu says goodbye")
}

func showInputLine() {
	fmt.Print(">   ")
}

func showContinuationLine() {
	fmt.Print("... ")
}

func showOutput(output string) {
	fmt.Println(output)
}

func handleLineOfText(line string) (bool, bool, string) {
	p := parser.New(parser.ReplMode)
	ok, isBreak, e := p.Parse(line)
	return ok, isBreak, e.String()
}
