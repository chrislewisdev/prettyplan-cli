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

// import { parseChangeSymbol } from '../src/ts/parse';

// test('parse change symbol', function() {
//     expect(parseChangeSymbol('+')).toBe('create');
//     expect(parseChangeSymbol('-')).toBe('destroy');
//     expect(parseChangeSymbol('~')).toBe('update');
//     expect(parseChangeSymbol('<=')).toBe('read');
//     expect(parseChangeSymbol('-/+')).toBe('recreate');
//     expect(parseChangeSymbol('gibberish')).toBe('unknown');
// });
