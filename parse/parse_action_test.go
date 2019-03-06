package parse

import "testing"

func TestParseCreateAction(t *testing.T) {
	action := parseAction(`
		+ module.alb.aws_alb_listener.default_https
			ssl_policy:                                             "old" => "new"
	`)

	expectString(t, "default_https", action.Id.Name)
	expectString(t, ChangeTypeCreate, action.Type)
	expectInt(t, 1, len(action.Diffs))
}

func TestParseReadAction(t *testing.T) {
	action := parseAction(`
		<= module.alb.aws_alb_listener.default_https
			ssl_policy:                                             "old" => "new"
	`)

	expectString(t, "default_https", action.Id.Name)
	expectString(t, ChangeTypeRead, action.Type)
	expectInt(t, 1, len(action.Diffs))
}

func TestParseUpdateAction(t *testing.T) {
	action := parseAction(`
		~ module.api.aws_alb_listener_rule.default
			condition.2636223071.field:                             "path-pattern" => ""
			condition.2636223071.field:                             "path-pattern" => ""
			condition.2636223071.field:                             "path-pattern" => ""
	`)

	expectString(t, "default", action.Id.Name)
	expectString(t, ChangeTypeUpdate, action.Type)
	expectInt(t, 3, len(action.Diffs))
}

func TestParseRecreateAction(t *testing.T) {
	action := parseAction(`
		-/+ module.api.aws_alb_listener_rule.default
			condition.2636223071.field:                             "path-pattern" => ""
			condition.2636223071.field:                             "path-pattern" => ""
			condition.2636223071.field:                             "path-pattern" => ""
	`)

	expectString(t, "default", action.Id.Name)
	expectString(t, ChangeTypeRecreate, action.Type)
	expectInt(t, 3, len(action.Diffs))
}

func TestParseDestroyAction(t *testing.T) {
	action := parseAction(`
		- module.alb.aws_alb_listener.default_https
	`)

	expectString(t, "default_https", action.Id.Name)
	expectString(t, ChangeTypeDestroy, action.Type)
	expectInt(t, 0, len(action.Diffs))
}
