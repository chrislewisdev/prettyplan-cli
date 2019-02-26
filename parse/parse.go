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
	// return ResourceId{}
	idSegments := strings.Split(resourceId, ".")

	resourceName := idSegments[len(idSegments)-1]
	resourceType := idSegments[len(idSegments)-2]
	resourcePrefixes := idSegments[0 : len(idSegments)-2]

	return ResourceId{
		Name:     resourceName,
		Type:     resourceType,
		Prefixes: resourcePrefixes}
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
