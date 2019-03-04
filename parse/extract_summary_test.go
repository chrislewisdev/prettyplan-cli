package parse

import "testing"

func TestExtractSingleLineChangeSummary(t *testing.T) {
	summary := extractPlanSummary("Terraform will perform the following actions:<summary>")

	expectString(t, "<summary>", summary)
}

func TestExtractMultiLineChangeSummary(t *testing.T) {
	summary := extractPlanSummary(`
		Text preceding the terraform plan

		Terraform will perform the following actions:

		<summary>`)

	expectString(t, "\n\n		<summary>", summary)
}

func TestExtractPlainStringSummary(t *testing.T) {
	summary := extractPlanSummary("<summary>")

	expectString(t, "<summary>", summary)
}
