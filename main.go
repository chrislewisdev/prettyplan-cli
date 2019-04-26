//go:generate packr2
package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"time"

	"github.com/chrislewisdev/prettyplan-cli/parse"
	"github.com/chrislewisdev/prettyplan-cli/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/browser"
)

type report struct {
	Version    string
	Plan       parse.Plan
	RawPlan    string
	Styles     template.CSS
	Scripts    template.JS
	ReportTime string
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	openFileFlag := flag.Bool("open", false, "To open the HTML report once generated")
	flag.Parse()

	templates := packr.New("Templates", "./templates")

	styles, err := templates.FindString("style.css")
	panicIfError(err)

	scripts, err := templates.FindString("scripts.js")
	panicIfError(err)

	html, err := templates.FindString("prettyplan.html")
	panicIfError(err)

	templateFunctions := template.FuncMap{
		"prettify": render.Prettify,
	}

	reportTemplate, err := template.New("prettyplan.html").Funcs(templateFunctions).Parse(html)
	panicIfError(err)

	rawPlan, err := exec.Command("terraform", "plan", "-no-color").Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("terraform plan failed! Output:\n %v", string(exitError.Stderr))
		} else {
			fmt.Printf("terraform plan failed! Output:\n %v", err.Error())
		}
		return
	}

	plan := parse.Parse(string(rawPlan))

	outputFile, err := os.Create("prettyplan.html")
	panicIfError(err)

	err = reportTemplate.Execute(outputFile, report{
		Version:    "v1.1",
		Plan:       plan,
		RawPlan:    string(rawPlan),
		Styles:     template.CSS(styles),
		Scripts:    template.JS(scripts),
		ReportTime: time.Now().Format(time.ANSIC),
	})
	panicIfError(err)

	outputFile.Close()

	if *openFileFlag {
		browser.OpenFile("prettyplan.html")
	}
}
