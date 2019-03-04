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
