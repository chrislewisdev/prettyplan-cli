package parse

import "testing"

func TestNewAndOldValueDiffsWithQuoteFormatting(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("property_name: \"old_value\" => \"new_value\"")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, "old_value", diffs[0].OldValue)
	expectString(t, "new_value", diffs[0].NewValue)
}

func TestNewAndOldValueDiffsWithEmptyQuotes(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("property_name: \"\" => \"new_value\"")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, "", diffs[0].OldValue)
	expectString(t, "new_value", diffs[0].NewValue)
}

func TestNewAndOldValueDiffsWithComputedValues(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("property_name: \"old_value\" => <computed>")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, "old_value", diffs[0].OldValue)
	expectString(t, "<computed>", diffs[0].NewValue)
}

func TestNewAndOldValueDiffsWhitespaceHandling(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("   property_name   : \" old_value \" => \"new_value \"")

	expectInt(t, 1, len(diffs))
	expectString(t, "property_name", diffs[0].Property)
	expectString(t, " old_value ", diffs[0].OldValue)
	expectString(t, "new_value ", diffs[0].NewValue)
}

func TestNewAndOldValueDiffsMultiLine(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("property1: \"old1\" => \"new1\"\n property2: \"old2\" => \"new2\"")

	expectInt(t, 2, len(diffs))
	expectString(t, "property1", diffs[0].Property)
	expectString(t, "old1", diffs[0].OldValue)
	expectString(t, "new1", diffs[0].NewValue)
	expectString(t, "property2", diffs[1].Property)
	expectString(t, "old2", diffs[1].OldValue)
	expectString(t, "new2", diffs[1].NewValue)
}

func TestParseNewAndOldValueDiffsWithBigJsonDocument(t *testing.T) {
	diffs := parseNewAndOldValueDiffs(`
    policy:                                     "{\\r\\n    \\"Version\\": \\"2012-10-17\\",\\r\\n    \\"Statement\\": [\\r\\n        {\\r\\n            \\"Action\\": [\\r\\n                \\"sqs:*\\",\\r\\n                \\"sns:*\\"\\r\\n\\r\\n            ],\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": \\"*\\"\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": [\\r\\n                \\"*\\",\\r\\n                \\"*\\"\\r\\n            ],\\r\\n            \\"Action\\": [\\r\\n                \\"logs:CreateLogGroup\\",\\r\\n                \\"logs:CreateLogStream\\",\\r\\n                \\"logs:PutLogEvents\\"\\r\\n            ]\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": [\\r\\n                \\"*\\"\\r\\n            ],\\r\\n            \\"Action\\": [\\r\\n
        \\"s3:PutObject\\",\\r\\n                \\"s3:GetObject\\",\\r\\n                \\"s3:GetObjectVersion\\"\\r\\n            ]\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Action\\": [\\r\\n                \\"cloudformation:*\\"\\r\\n            ],\\r\\n            \\"Resource\\": \\"*\\"\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n"
    => "{\\r\\n    \\"Version\\": \\"2012-10-17\\",\\r\\n    \\"Statement\\": [\\r\\n        {\\r\\n            \\"Action\\": [\\r\\n                \\"sqs:*\\",\\r\\n                \\"sns:*\\"\\r\\n\\r\\n            ],\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": \\"*\\"\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": [\\r\\n                \\"*\\",\\r\\n                \\"*\\"\\r\\n            ],\\r\\n            \\"Action\\": [\\r\\n                \\"logs:CreateLogGroup\\",\\r\\n                \\"logs:CreateLogStream\\",\\r\\n                \\"logs:PutLogEvents\\"\\r\\n            ]\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Resource\\": [\\r\\n                \\"*\\"\\r\\n            ],\\r\\n            \\"Action\\": [\\r\\n
        \\"s3:PutObject\\",\\r\\n                \\"s3:GetObject\\",\\r\\n                \\"s3:GetObjectVersion\\"\\r\\n            ]\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n            \\"Action\\": [\\r\\n                \\"cloudformation:*\\"\\r\\n            ],\\r\\n            \\"Resource\\": \\"*\\"\\r\\n        },\\r\\n        {\\r\\n            \\"Effect\\": \\"Allow\\",\\r\\n"
	`)

	expectInt(t, 1, len(diffs))
	expectString(t, "policy", diffs[0].Property)
}

func TestForceNewResourceIsFalseForNormalDiffs(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("property_name: \"old_value\" => \"new_value\"")

	expectInt(t, 1, len(diffs))
	expectBool(t, false, diffs[0].ForcesNewResource)
}

func TestForceNewResourceIsTrueWhenIndicated(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("property_name: \"old_value\" => \"new_value\" (forces new resource)")

	expectInt(t, 1, len(diffs))
	expectBool(t, true, diffs[0].ForcesNewResource)
}

func TestForceNewResourceIsTrueWhenIndicatedOnComputedValues(t *testing.T) {
	diffs := parseNewAndOldValueDiffs("property_name: \"old_value\" => <computed> (forces new resource)")

	expectInt(t, 1, len(diffs))
	expectBool(t, true, diffs[0].ForcesNewResource)
}
