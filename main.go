package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type report struct {
	Version string
	Output  string
	Styles  template.CSS
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// terraformOutput, err := exec.Command("terraform", "plan", "-no-color").Output()
	// if err != nil {
	// 	if exitError, ok := err.(*exec.ExitError); ok {
	// 		fmt.Printf("terraform plan failed! Output:\n %v", string(exitError.Stderr))
	// 	} else {
	// 		fmt.Printf("terraform plan failed! Output:\n %v", err.Error())
	// 	}
	// 	return
	// }

	styles, err := ioutil.ReadFile("templates/style.css")
	panicIfError(err)

	reportTemplate, err := template.ParseFiles("templates/prettyplan.html")
	panicIfError(err)

	outputFile, err := os.Create("prettyplan.html")
	panicIfError(err)

	err = reportTemplate.Execute(outputFile, report{Version: "v1.2", Output: "Hello World", Styles: template.CSS(styles)})
	panicIfError(err)

	outputFile.Close()
}
