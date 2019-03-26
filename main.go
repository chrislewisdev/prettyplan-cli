//go:generate packr2
package main

import (
	"html/template"
	"os"

	"github.com/chrislewisdev/prettyplan-cli/parse"
	"github.com/chrislewisdev/prettyplan-cli/render"
	"github.com/gobuffalo/packr/v2"
)

type report struct {
	Version string
	Plan    parse.Plan
	RawPlan string
	Styles  template.CSS
	Scripts template.JS
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
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

	// rawPlan, err := exec.Command("terraform", "plan", "-no-color").Output()
	// if err != nil {
	// 	if exitError, ok := err.(*exec.ExitError); ok {
	// 		fmt.Printf("terraform plan failed! Output:\n %v", string(exitError.Stderr))
	// 	} else {
	// 		fmt.Printf("terraform plan failed! Output:\n %v", err.Error())
	// 	}
	// 	return
	// }

	rawPlan := `
	Warning: module.thing.aws_ecs_service.service: here is an example warning
	Warning: aws_ecs_service.service: here is another example warning

	Refreshing Terraform state in-memory prior to plan...
	The refreshed state will be used to calculate this plan, but will not be
	persisted to local or remote state storage.

	aws_alb_target_group.sample_app: Refreshing state... (ID: arn:aws:elasticloadbalancing:us-east-1:...up/sample-app/d5eedf0680cc9834)
	aws_iam_role.service_role: Refreshing state... (ID: SampleApp)
	aws_cloudwatch_log_group.sample_app: Refreshing state... (ID: sample-app)
	aws_ecr_repository.sample_app: Refreshing state... (ID: sample-app)
	aws_iam_role_policy.service_role_policy: Refreshing state... (ID: SampleApp:SampleApp)
	null_resource.promote_images: Refreshing state... (ID: 1236159896537553123)
	aws_ecs_task_definition.sample_app: Refreshing state... (ID: sample-app)
	aws_alb_listener_rule.routing: Refreshing state... (ID: arn:aws:elasticloadbalancing:us-east-1:...94bc/2825bddee1920172/ec8bc47bb5409ead)
	aws_ecs_service.sample_app: Refreshing state... (ID: arn:aws:ecs:us-east-1:123123123123:service/sample-app)

	------------------------------------------------------------------------

	An execution plan has been generated and is shown below.
	Resource actions are indicated with the following symbols:
		~ update in-place
	-/+ destroy and then create replacement
		<= read (data resources)

	Terraform will perform the following actions:

		<= data.external.ecr_image_digests
			id:                       <computed>
			program.#:                "1"
			program.0:                "extract-image-digests"
			result.%:                 <computed>

		- aws_ecs_service.sample_app
			task_definition:          "arn:aws:ecs:us-east-1:123123123123:task-definition/sample-app:186" => "${ aws_ecs_task_definition.sample_app.arn }"

	-/+ aws_ecs_task_definition.sample_app (new resource required)
			id:                       "sample-app" => <computed> (forces new resource)
			arn:                      "arn:aws:ecs:us-east-1:123123123123:task-definition/sample-app:186" => <computed>
			container_definitions:    "[{\"cpu\":1,\"environment\":[],\"essential\":true,\"image\":\"123123123123.dkr.ecr.us-east-1.amazonaws.com/sample-app@sha256:18979dcf521de65f736585d30b58e8085ecc44560fa8c530ad1eb17fecad1cab\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"sample-app\",\"awslogs-region\":\"us-east-1\",\"awslogs-stream-prefix\":\"sample-app\"}},\"memory\":256,\"mountPoints\":[],\"name\":\"sample-app\",\"portMappings\":[{\"containerPort\":8443,\"hostPort\":0,\"protocol\":\"tcp\"}],\"volumesFrom\":[]}]" => "[\n  {\n    \"name\": \"sample-app\",\n    \"image\": \"${ aws_ecr_repository.sample_app.repository_url }@${ data.external.ecr_image_digests.result[\"sample-app\"] }\",\n    \"cpu\": 1,\n    \"memory\": 256,\n    \"essential\": true,\n    \"logConfiguration\": {\n      \"logDriver\": \"awslogs\",\n      \"options\": {\n        \"awslogs-group\": \"${ aws_cloudwatch_log_group.sample_app.name }\",\n        \"awslogs-region\": \"${ var.target_aws_region }\",\n        \"awslogs-stream-prefix\": \"sample-app\"\n      }\n    },\n    \"portMappings\": [\n      {\n        \"containerPort\": 8443,\n        \"hostPort\": 0\n      }\n    ]\n  }\n]\n" (forces new resource)
			family:                   "sample-app" => "sample-app"
			network_mode:             "" => <computed>
			revision:                 "186" => <computed>
			task_role_arn:            "arn:aws:iam::123123123123:role/SampleApp" => "arn:aws:iam::123123123123:role/SampleApp"

	-/+ null_resource.promote_images (new resource required)
			id:                       "1236159896537553123" => <computed> (forces new resource)
			triggers.%:               "1" => "1.2"
			triggers.deploy_job_hash: "6c37ac7175bdf35e24a2f2755addd238" => "1a0bd86fc5831ee66858f2e159efa547" (forces new resource)

	Plan: 2 to add, 1 to change, 2 to destroy.

	------------------------------------------------------------------------

	This plan was saved to: terraform.plan

	To perform exactly these actions, run the following command to apply:
		terraform apply "terraform.plan"
		`

	plan := parse.Parse(string(rawPlan))

	outputFile, err := os.Create("prettyplan.html")
	panicIfError(err)

	err = reportTemplate.Execute(outputFile, report{
		Version: "v1.2",
		Plan:    plan,
		RawPlan: string(rawPlan),
		Styles:  template.CSS(styles),
		Scripts: template.JS(scripts),
	})
	panicIfError(err)

	outputFile.Close()
}
