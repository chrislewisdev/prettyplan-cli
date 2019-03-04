package parse

import "strings"

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

	warnings = append(warnings, Warning{
		Id: ResourceId{
			Name:     "gateway",
			Type:     "transit_gateway",
			Prefixes: nil},
		Detail: "warning"})

	return warnings
}
