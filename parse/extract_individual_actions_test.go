package parse

import "testing"

func TestExtractActionsWithSummaryAtEnd(t *testing.T) {
	actions := extractIndividualActions(`
		+ module.alb.aws_alb_listener.default_https
			ssl_policy:                                             "old" => "new"
	
		~ module.api.aws_alb_listener_rule.default
			condition.2636223071.field:                             "path-pattern" => ""
			condition.2636223071.field:                             "path-pattern" => ""
			condition.2636223071.field:                             "path-pattern" => ""
	
		- module.api.aws_alb_target_group.default
			health_check.0.path:                                    "/healthcheck/old" => "/healthcheck/new"
	
		-/+ module.service_a.aws_ecs_service.default
			task_definition:                                        "service-a:185" => "service-a:179"
	
		<= module.service_b.aws_ecs_service.default
			task_definition:                                        "service-b:171" => "service-b:165"
		Plan: 2 to add, 1 to change, 2 to destroy.`)

	expectInt(t, 5, len(actions))
}

func TestExtractActionsWithNoSummary(t *testing.T) {
	actions := extractIndividualActions(`
		+ module.alb.aws_alb_listener.default_https
			ssl_policy:                                             "old" => "new"

		~ module.api.aws_alb_listener_rule.default
			condition.2636223071.field:                             "path-pattern" => ""
	`)

	expectInt(t, 2, len(actions))
}

func TestExtractActionsWithPrefixText(t *testing.T) {
	actions := extractIndividualActions(`
		this text here should not be detected part of the change
		neither should this
		-------------------------------------------

		+ module.alb.aws_alb_listener.default_https
			ssl_policy:                                             "old" => "new"

		~ module.api.aws_alb_listener_rule.default
			condition.2636223071.field:                             "path-pattern" => ""
	`)

	expectInt(t, 2, len(actions))
}

func TestExtractEmptyActions(t *testing.T) {
	actions := extractIndividualActions("random text that has no actions")

	expectInt(t, 0, len(actions))
}
