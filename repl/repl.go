package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rokob/kudu/interpreter"
	"github.com/rokob/kudu/parser"
)

// Run - run the repl
func Run() {
	interpreter := interpreter.New()
	scanner := bufio.NewScanner(os.Stdin)
	showTagline()
	showInputLine()
	var currentInput string
	for scanner.Scan() {
		lastInput := scanner.Text()
		if lastInput == ".bye" {
			break
		}
		currentInput += lastInput
		ok, isBreak, output := handleLineOfText(currentInput, interpreter)
		if !ok && !isBreak {
			okLast, isBreakLast, _ := handleLineOfText(lastInput, interpreter)
			if !okLast && isBreakLast {
				currentInput = ""
				showInputLine()
			} else {
				currentInput += "\n"
				showContinuationLine()
			}
			continue
		}
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
	showExit()
}

func showTagline() {
	fmt.Println("🦌🦌🦌 Kudu v0.1 - The Kuduest Language 🦌🦌🦌")
	fmt.Println("Use $ to exit an expression")
	fmt.Println("^D or .bye to quit")
}

func showExit() {
	fmt.Println("\n🦌🦌🦌 Kudu says goodbye")
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

func handleLineOfTextToAST(line string) (bool, bool, string) {
	p := parser.New(parser.ReplMode)
	ok, isBreak, es := p.Parse(line)
	if ok {
		var b []byte
		var err error
		if len(es) == 1 {
			b, err = json.MarshalIndent(es[0], "", "  ")
		} else {
			b, err = json.MarshalIndent(es, "", "  ")
		}
		if err != nil {
			return false, false, err.Error()
		}
		return ok, isBreak, string(b)
	}
	return ok, isBreak, ""
}

func handleLineOfText(line string, interpreter *interpreter.Interpreter) (bool, bool, string) {
	p := parser.New(parser.ReplMode)
	ok, isBreak, es := p.Parse(line)
	if ok {
		results := make([]string, len(es))
		for i, e := range es {
			results[i] = interpreter.HandleExpression(e).String()
		}
		return true, false, strings.Join(results, "\n")
	}
	return ok, isBreak, ""
}
