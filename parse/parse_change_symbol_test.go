package parse

import "testing"

func TestParseChangeSymbol(t *testing.T) {
	expectString(t, ChangeTypeCreate, parseChangeSymbol("+"))
	expectString(t, ChangeTypeDestroy, parseChangeSymbol("-"))
	expectString(t, ChangeTypeUpdate, parseChangeSymbol("~"))
	expectString(t, ChangeTypeRead, parseChangeSymbol("<="))
	expectString(t, ChangeTypeRecreate, parseChangeSymbol("-/+"))
	expectString(t, ChangeTypeUnknown, parseChangeSymbol("gibberish"))
}
