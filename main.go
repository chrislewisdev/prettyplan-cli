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

func main() {
	openFileFlag := flag.Bool("open", false, "To open the HTML report once generated")
	fileNameFlag := flag.String("filename", "prettyplan-output.html", "Specify a filename other than prettyplan-output.html")
	flag.Parse()

	assets := getAssets()

	var plan *planInfo
	var err error
	if plan, err = getPlan(); err != nil {
		fmt.Println(err.Error())
		return
	}

	writeReport(*fileNameFlag, assets, plan)

	if *openFileFlag {
		browser.OpenFile(*fileNameFlag)
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type templateAssets struct {
	Styles   template.CSS
	Scripts  template.JS
	Template *template.Template
}

func getAssets() *templateAssets {
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

	return &templateAssets{
		Styles:   template.CSS(styles),
		Scripts:  template.JS(scripts),
		Template: reportTemplate,
	}
}

type planInfo struct {
	Raw    string
	Parsed parse.Plan
}

func getPlan() (*planInfo, error) {
	rawPlan, err := exec.Command("terraform", "plan", "-no-color").Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("terraform plan failed! Output:\n %v", string(exitError.Stderr))
		}
		return nil, fmt.Errorf("terraform plan failed! Output:\n %v", err.Error())
	}

	plan := parse.Parse(string(rawPlan))

	return &planInfo{Raw: string(rawPlan), Parsed: plan}, nil
}

type reportInfo struct {
	Version    string
	Plan       parse.Plan
	RawPlan    string
	Styles     template.CSS
	Scripts    template.JS
	ReportTime string
}

func writeReport(fileName string, assets *templateAssets, plan *planInfo) {
	outputFile, err := os.Create(fileName)
	panicIfError(err)

	err = assets.Template.Execute(outputFile, reportInfo{
		Version:    "v1.1",
		Plan:       plan.Parsed,
		RawPlan:    plan.Raw,
		Styles:     assets.Styles,
		Scripts:    assets.Scripts,
		ReportTime: time.Now().Format(time.ANSIC),
	})
	panicIfError(err)

	outputFile.Close()
}
