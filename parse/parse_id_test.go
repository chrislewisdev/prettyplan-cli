package parse

import (
	"testing"
)

func TestParseIdWithoutPrefixes(t *testing.T) {
	id := parseId("aws_route53_record.domain_name")

	expectString(t, "domain_name", id.Name)
	expectString(t, "aws_route53_record", id.Type)
	expectInt(t, 0, len(id.Prefixes))
}

func TestParseIdWithPrefixes(t *testing.T) {
	id := parseId("module.api.aws_ecs_service.api_service")

	expectString(t, "api_service", id.Name)
	expectString(t, "aws_ecs_service", id.Type)
	expectStrings(t, []string{"module", "api"}, id.Prefixes)
}

func TestParseIdWithNameOnly(t *testing.T) {
	id := parseId("api_service")

	expectString(t, "api_service", id.Name)
	expectString(t, "", id.Type)
	expectInt(t, 0, len(id.Prefixes))
}
