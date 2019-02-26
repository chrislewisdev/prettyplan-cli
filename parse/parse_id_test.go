package parse

import "testing"

func TestParseIdNoPrefixes(t *testing.T) {
	id := parseId("aws_route53_record.domain_name")

	if id.Name != "domain_name" {
		t.Error("Expected resource name to be 'domain_name', got: ", id.Name)
	}
	if id.Type != "aws_route53_record" {
		t.Error("Expected resource type to be 'aws_route53_record', got: ", id.Type)
	}
	if len(id.Prefixes) != 0 {
		t.Error("Expected no prefixes")
	}
}

// import { parseId } from '../src/ts/parse'

// test('parse id - no prefixes', function() {
//     const id = parseId('aws_route53_record.domain_name');

//     expect(id.name).toBe('domain_name');
//     expect(id.type).toBe('aws_route53_record');
//     expect(id.prefixes).toEqual([]);
// });
// test('parse id - with prefixes', function() {
//     const id = parseId('module.api.aws_ecs_service.api_service');

//     expect(id.name).toBe('api_service');
//     expect(id.type).toBe('aws_ecs_service');
//     expect(id.prefixes).toEqual(['module', 'api']);
// });
// test('parse id - name only', function() {
//     const id = parseId('api_service');

//     expect(id.name).toBe('api_service');
//     expect(id.type).toBeNull();
//     expect(id.prefixes).toEqual([]);
// });
