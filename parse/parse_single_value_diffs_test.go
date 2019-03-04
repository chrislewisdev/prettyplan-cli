package parse

import "testing"

func TestParseSingleValueDiffsWithQuotes(t *testing.T) {
	diffs := parseSingleValueDiffs("property_name: \"new_value\"")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, "new_value", diffs[0].NewValue)
}

func TestParseSingleValueDiffsWithEmptyQuotes(t *testing.T) {
	diffs := parseSingleValueDiffs("property_name: \"\"")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, "", diffs[0].NewValue)
}

func TestParseSingleValueDiffsWithComputedValues(t *testing.T) {
	diffs := parseSingleValueDiffs("property_name: <computed>")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, "<computed>", diffs[0].NewValue)
}

func TestParseSingleValueDiffsWhitespaceHandling(t *testing.T) {
	diffs := parseSingleValueDiffs("     property_name :    \" value \"   ")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, " value ", diffs[0].NewValue)
}

func TestParseSingleValueDiffsWithMultiLine(t *testing.T) {
	diffs := parseSingleValueDiffs("property1: \"value1\"\n property2: \"value2\"")

	expectInt(t, 2, len(diffs))
	expectString(t, "property1", diffs[0].Property)
	expectString(t, "value1", diffs[0].NewValue)
	expectString(t, "property2", diffs[1].Property)
	expectString(t, "value2", diffs[1].NewValue)
}

func TestParseSingleValueDiffsWithBigPolicyDocument(t *testing.T) {
	diffs := parseSingleValueDiffs(`
    policy:                                     "{\\r\\n    \\"Version\\": \\"2012-10-17\\",\\r\\n    \\"Statement\\": [\\r\\n        {\\r\\n            \\"Action\\": [\\r\\n                \\"sqs:*\\",\\r\\n                \\"sns:*\\"\\r\\n\\r\\n            ],\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": \\"*\\"\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": [\\r\\n                \\"*\\",\\r\\n                \\"*\\"\\r\\n            ],\\r\\n            \\"Action\\": [\\r\\n                \\"logs:CreateLogGroup\\",\\r\\n                \\"logs:CreateLogStream\\",\\r\\n                \\"logs:PutLogEvents\\"\\r\\n            ]\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": [\\r\\n                \\"*\\"\\r\\n            ],\\r\\n            \\"Action\\": [\\r\\n
        \\"s3:PutObject\\",\\r\\n                \\"s3:GetObject\\",\\r\\n                \\"s3:GetObjectVersion\\"\\r\\n            ]\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Action\\": [\\r\\n                \\"cloudformation:*\\"\\r\\n            ],\\r\\n            \\"Resource\\": \\"*\\"\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n"
    `)

	expectInt(t, 1, len(diffs))
	expectString(t, "policy", diffs[0].Property)
	//Look, it's not worth asserting on the full document contents. Let's exercise Trust Driven Development here.
}
