package parse

import "testing"

func testParseSingleWarning(t *testing.T) {
	warnings := parseWarnings("Warning: resource_name: <warning detail>")

	if len(warnings) != 1 {
		t.Error("Expected 1 warning, got ", len(warnings))
	}
	if warnings[0].Id.Name != "resource_name" {
		t.Error("Expected resource name to be 'resource_name, got ", warnings[0].Id.Name)
	}
}

// test('parse warnings - single warning', function() {
//     const warnings = parseWarnings('Warning: resource_name: <warning detail>');

//     expect(warnings).toHaveLength(1);
//     expect(warnings[0].id.name).toBe('resource_name:');
//     expect(warnings[0].detail).toBe(' <warning detail>');
// });

// test('parse warnings - multiple warning', function() {
//     const warnings = parseWarnings('Warning: r1: w1\nWarning: r2: w2\nWarning: r3: w3');

//     expect(warnings).toHaveLength(3);

//     expect(warnings[0].id.name).toBe('r1:');
//     expect(warnings[0].detail).toBe(' w1');

//     expect(warnings[1].id.name).toBe('r2:');
//     expect(warnings[1].detail).toBe(' w2');

//     expect(warnings[2].id.name).toBe('r3:');
//     expect(warnings[2].detail).toBe(' w3');
// });

// test('parse warnings - no warnings', function() {
//     const warnings = parseWarnings('Here are some things that are NOT warnings');

//     expect(warnings).toHaveLength(0);
// });
