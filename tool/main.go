package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
)

const (
	templateFileName = `tool/problem.tmpl`

	mainFileName = `main.go`
	mainMarker   = `// TEMPLATE_MARKER //`
)

func main() {
	ep := easyPuzzle{}

	err := survey.Ask([]*survey.Question{{
		Name:   "number",
		Prompt: &survey.Input{Message: "What's the puzzle number?"},
	}, {
		Name:   "prompt",
		Prompt: &survey.Input{Message: "What's the prompt?"},
	}}, &ep)
	if err != nil {
		log.Fatal(err)
	}

	newEasyPuzzle(ep)
}

type easyPuzzle struct {
	Prompt string
	Number int
}

func newEasyPuzzle(
	ep easyPuzzle,
) {
	if ep.Number > 100 {
		log.Printf("We will not generate puzzles for puzzles beyond the first 100.\n")
		return
	} else if ep.Number <= 0 {
		log.Printf("got an invalid number: %d\n", ep.Number)
		return
	}

	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("easy/puzzle%d.go", ep.Number)
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.ExecuteTemplate(f, `easyPuzzleFile`, ep)
	if err != nil {
		log.Fatal(err)
	}

	mainF, err := os.Open(mainFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer mainF.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, `caseForMainSolver`, ep)
	if err != nil {
		log.Fatal(err)
	}
}
