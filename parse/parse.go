package parse

import (
	"regexp"
	"strings"
)

type ResourceId struct {
	Name     string
	Type     string
	Prefixes []string
}
type Warning struct {
	Id     ResourceId
	Detail string
}

const (
	ChangeTypeCreate   string = "create"
	ChangeTypeDestroy  string = "destroy"
	ChangeTypeUpdate   string = "update"
	ChangeTypeRead     string = "read"
	ChangeTypeRecreate string = "recreate"
	ChangeTypeUnknown  string = "unknown"
)

type Diff struct {
	Property          string
	OldValue          string
	NewValue          string
	ForcesNewResource bool
}

func parseId(resourceId string) ResourceId {
	idSegments := strings.Split(resourceId, ".")

	resource := ResourceId{}

	resource.Name = idSegments[len(idSegments)-1]
	if len(idSegments) > 1 {
		resource.Type = idSegments[len(idSegments)-2]
	}
	if len(idSegments) > 2 {
		resource.Prefixes = idSegments[0 : len(idSegments)-2]
	}

	return resource
}

func parseWarnings(plan string) []Warning {
	warnings := make([]Warning, 0)
	r := regexp.MustCompile("Warning: (.*):(.*)")

	for _, match := range r.FindAllStringSubmatch(plan, -1) {
		warnings = append(warnings, Warning{
			Id:     parseId(match[1]),
			Detail: match[2]})
	}

	return warnings
}

func extractPlanSummary(plan string) string {
	splits := strings.SplitAfter(plan, "Terraform will perform the following actions:")
	return splits[len(splits)-1]
}

func extractIndividualActions(actionSummary string) []string {
	//In JS, a negative lookahead was used to accurately capture each distinct action and its diffs
	//But in Go we can't use lookaheads, so instead we look for the start of each action and take the substrings between them
	r := regexp.MustCompile(`([~+-]|-\/\+|<=) .*`)
	matches := r.FindAllStringIndex(actionSummary, -1)
	actions := make([]string, 0)

	for i := 0; i < len(matches)-1; i++ {
		actions = append(actions, actionSummary[matches[i][0]:matches[i+1][0]])
	}
	if len(matches) >= 1 {
		actions = append(actions, actionSummary[matches[len(matches)-1][0]:])
	}

	return actions
}

func parseChangeSymbol(changeSymbol string) string {
	switch changeSymbol {
	case "-":
		return ChangeTypeDestroy
	case "+":
		return ChangeTypeCreate
	case "~":
		return ChangeTypeUpdate
	case "-/+":
		return ChangeTypeRecreate
	case "<=":
		return ChangeTypeRead
	default:
		return ChangeTypeUnknown
	}
}

func takeFirstNonEmptyString(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return values[len(values)-1]
}

func parseSingleValueDiffs(action string) []Diff {
	diffs := make([]Diff, 0)
	r := regexp.MustCompile(`\s*(.*?): *(?:(<computed>)|"(|[\S\s]*?[^\\])")`)

	for _, match := range r.FindAllStringSubmatch(action, -1) {
		diffs = append(diffs, Diff{
			Property: strings.TrimSpace(match[1]),
			NewValue: takeFirstNonEmptyString(match[2], match[3])})
	}

	return diffs
}

func parseNewAndOldValueDiffs(action string) []Diff {
	diffs := make([]Diff, 0)

	r := regexp.MustCompile(`\s*(.*?): *(?:"(|[\S\s]*?[^\\])")[\S\s]*?=> *(?:(<computed>)|"(|[\S\s]*?[^\\])")( \(forces new resource\))?`)

	for _, match := range r.FindAllStringSubmatch(action, -1) {
		diffs = append(diffs, Diff{
			Property:          strings.TrimSpace(match[1]),
			OldValue:          match[2],
			NewValue:          takeFirstNonEmptyString(match[3], match[4]),
			ForcesNewResource: match[5] != ""})
	}

	return diffs
}
