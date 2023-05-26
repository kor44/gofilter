
//line lexer.rl:1
package gofilter

import "fmt"


//line lexer.rl:109



//line lexer.go:13
var _filterLexerImpl_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 15, 1, 16, 
	1, 17, 1, 18, 1, 19, 1, 20, 
	1, 21, 1, 22, 1, 23, 1, 24, 
	1, 25, 1, 26, 1, 27, 1, 28, 
	1, 29, 1, 30, 1, 31, 1, 32, 
	2, 2, 3, 2, 2, 4, 2, 2, 
	5, 2, 2, 6, 2, 2, 7, 2, 
	2, 8, 2, 2, 9, 2, 2, 10, 
	2, 2, 11, 2, 2, 12, 2, 2, 
	13, 2, 2, 14, 
}

var _filterLexerImpl_key_offsets []int16 = []int16{
	0, 0, 2, 2, 3, 5, 6, 7, 
	39, 40, 50, 59, 61, 62, 63, 73, 
	83, 93, 103, 113, 123, 133, 143, 153, 
	163, 174, 185, 195, 205, 215, 225, 235, 
	245, 256, 266, 
}

var _filterLexerImpl_trans_keys []byte = []byte{
	34, 92, 38, 48, 57, 61, 124, 32, 
	33, 34, 38, 40, 41, 43, 44, 60, 
	61, 62, 95, 97, 99, 101, 103, 108, 
	109, 110, 111, 124, 126, 9, 12, 45, 
	46, 48, 58, 65, 90, 98, 122, 61, 
	43, 95, 45, 46, 48, 58, 65, 90, 
	97, 122, 43, 47, 95, 45, 58, 65, 
	90, 97, 122, 48, 57, 61, 61, 43, 
	47, 95, 110, 45, 58, 65, 90, 97, 
	122, 43, 47, 95, 100, 45, 58, 65, 
	90, 97, 122, 43, 47, 95, 111, 45, 
	58, 65, 90, 97, 122, 43, 47, 95, 
	110, 45, 58, 65, 90, 97, 122, 43, 
	47, 95, 116, 45, 58, 65, 90, 97, 
	122, 43, 47, 95, 97, 45, 58, 65, 
	90, 98, 122, 43, 47, 95, 105, 45, 
	58, 65, 90, 97, 122, 43, 47, 95, 
	110, 45, 58, 65, 90, 97, 122, 43, 
	47, 95, 115, 45, 58, 65, 90, 97, 
	122, 43, 47, 95, 113, 45, 58, 65, 
	90, 97, 122, 43, 47, 95, 101, 116, 
	45, 58, 65, 90, 97, 122, 43, 47, 
	95, 101, 116, 45, 58, 65, 90, 97, 
	122, 43, 47, 95, 97, 45, 58, 65, 
	90, 98, 122, 43, 47, 95, 116, 45, 
	58, 65, 90, 97, 122, 43, 47, 95, 
	99, 45, 58, 65, 90, 97, 122, 43, 
	47, 95, 104, 45, 58, 65, 90, 97, 
	122, 43, 47, 95, 101, 45, 58, 65, 
	90, 97, 122, 43, 47, 95, 115, 45, 
	58, 65, 90, 97, 122, 43, 47, 95, 
	101, 111, 45, 58, 65, 90, 97, 122, 
	43, 47, 95, 116, 45, 58, 65, 90, 
	97, 122, 43, 47, 95, 114, 45, 58, 
	65, 90, 97, 122, 
}

var _filterLexerImpl_single_lengths []byte = []byte{
	0, 2, 0, 1, 0, 1, 1, 22, 
	1, 2, 3, 0, 1, 1, 4, 4, 
	4, 4, 4, 4, 4, 4, 4, 4, 
	5, 5, 4, 4, 4, 4, 4, 4, 
	5, 4, 4, 
}

var _filterLexerImpl_range_lengths []byte = []byte{
	0, 0, 0, 0, 1, 0, 0, 5, 
	0, 4, 3, 1, 0, 0, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 
}

var _filterLexerImpl_index_offsets []int16 = []int16{
	0, 0, 3, 4, 6, 8, 10, 12, 
	40, 42, 49, 56, 58, 60, 62, 70, 
	78, 86, 94, 102, 110, 118, 126, 134, 
	142, 151, 160, 168, 176, 184, 192, 200, 
	208, 217, 225, 
}

var _filterLexerImpl_indicies []byte = []byte{
	1, 2, 0, 0, 3, 4, 6, 5, 
	7, 4, 8, 4, 9, 10, 0, 11, 
	12, 13, 14, 15, 17, 18, 19, 16, 
	20, 21, 22, 23, 24, 25, 26, 27, 
	28, 29, 9, 16, 16, 16, 16, 4, 
	31, 30, 14, 14, 14, 14, 14, 14, 
	32, 14, 33, 16, 16, 16, 16, 5, 
	6, 34, 36, 35, 38, 37, 14, 33, 
	16, 39, 16, 16, 16, 32, 14, 33, 
	16, 40, 16, 16, 16, 32, 14, 33, 
	16, 41, 16, 16, 16, 32, 14, 33, 
	16, 42, 16, 16, 16, 32, 14, 33, 
	16, 43, 16, 16, 16, 32, 14, 33, 
	16, 44, 16, 16, 16, 32, 14, 33, 
	16, 45, 16, 16, 16, 32, 14, 33, 
	16, 46, 16, 16, 16, 32, 14, 33, 
	16, 47, 16, 16, 16, 32, 14, 33, 
	16, 48, 16, 16, 16, 32, 14, 33, 
	16, 49, 50, 16, 16, 16, 32, 14, 
	33, 16, 51, 52, 16, 16, 16, 32, 
	14, 33, 16, 53, 16, 16, 16, 32, 
	14, 33, 16, 54, 16, 16, 16, 32, 
	14, 33, 16, 55, 16, 16, 16, 32, 
	14, 33, 16, 56, 16, 16, 16, 32, 
	14, 33, 16, 57, 16, 16, 16, 32, 
	14, 33, 16, 58, 16, 16, 16, 32, 
	14, 33, 16, 59, 60, 16, 16, 16, 
	32, 14, 33, 16, 61, 16, 16, 16, 
	32, 14, 33, 16, 62, 16, 16, 16, 
	32, 
}

var _filterLexerImpl_trans_targs []byte = []byte{
	1, 7, 2, 7, 0, 7, 11, 7, 
	7, 7, 8, 3, 7, 7, 9, 7, 
	10, 12, 5, 13, 14, 16, 23, 24, 
	25, 26, 32, 34, 6, 7, 7, 7, 
	7, 4, 7, 7, 7, 7, 7, 15, 
	10, 17, 18, 19, 20, 21, 22, 10, 
	10, 10, 10, 10, 10, 27, 28, 29, 
	30, 31, 10, 10, 33, 10, 10, 
}

var _filterLexerImpl_trans_actions []byte = []byte{
	0, 25, 0, 21, 0, 39, 0, 11, 
	23, 27, 0, 0, 5, 7, 0, 9, 
	74, 0, 0, 0, 74, 74, 74, 74, 
	74, 74, 74, 74, 0, 19, 33, 13, 
	37, 0, 35, 31, 17, 29, 15, 74, 
	68, 74, 74, 74, 74, 74, 74, 59, 
	41, 50, 47, 56, 53, 74, 74, 74, 
	74, 74, 62, 44, 74, 65, 71, 
}

var _filterLexerImpl_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 
}

var _filterLexerImpl_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 3, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 
}

var _filterLexerImpl_eof_trans []int16 = []int16{
	0, 0, 0, 0, 6, 0, 0, 0, 
	31, 33, 6, 35, 36, 38, 33, 33, 
	33, 33, 33, 33, 33, 33, 33, 33, 
	33, 33, 33, 33, 33, 33, 33, 33, 
	33, 33, 33, 
}

const filterLexerImpl_start int = 7
const filterLexerImpl_first_final int = 7
const filterLexerImpl_error int = 0

const filterLexerImpl_en_main int = 7


//line lexer.rl:121



type filterLexerImpl struct {
	data []byte
	cs 	 int
  	p    int
	pe   int
	act  int
	ts   int
	te   int
	eof	 int
	result node
	err  string
}

func newLex(line []byte) *filterLexerImpl {
	lexer := filterLexerImpl{data: line}
	
//line lexer.go:203
	{
	( lexer.cs) = filterLexerImpl_start
	( lexer.ts) = 0
	( lexer.te) = 0
	( lexer.act) = 0
	}

//line lexer.rl:140
	lexer.pe = len(line)
	lexer.eof = len(line)
	return &lexer
}

//type FilterSymType struct {
//	Data []byte
//} // fake

func (lexer *filterLexerImpl) Lex(lval *filterSymType) int {
    token_kind := 0
	
//line lexer.go:224
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if ( lexer.p) == ( lexer.pe) {
		goto _test_eof
	}
	if ( lexer.cs) == 0 {
		goto _out
	}
_resume:
	_acts = int(_filterLexerImpl_from_state_actions[( lexer.cs)])
	_nacts = uint(_filterLexerImpl_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _filterLexerImpl_actions[_acts - 1] {
		case 1:
//line NONE:1
( lexer.ts) = ( lexer.p)

//line lexer.go:247
		}
	}

	_keys = int(_filterLexerImpl_key_offsets[( lexer.cs)])
	_trans = int(_filterLexerImpl_index_offsets[( lexer.cs)])

	_klen = int(_filterLexerImpl_single_lengths[( lexer.cs)])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case ( lexer.data)[( lexer.p)] < _filterLexerImpl_trans_keys[_mid]:
				_upper = _mid - 1
			case ( lexer.data)[( lexer.p)] > _filterLexerImpl_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_filterLexerImpl_range_lengths[( lexer.cs)])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case ( lexer.data)[( lexer.p)] < _filterLexerImpl_trans_keys[_mid]:
				_upper = _mid - 2
			case ( lexer.data)[( lexer.p)] > _filterLexerImpl_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_filterLexerImpl_indicies[_trans])
_eof_trans:
	( lexer.cs) = int(_filterLexerImpl_trans_targs[_trans])

	if _filterLexerImpl_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_filterLexerImpl_trans_actions[_trans])
	_nacts = uint(_filterLexerImpl_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _filterLexerImpl_actions[_acts-1] {
		case 2:
//line NONE:1
( lexer.te) = ( lexer.p)+1

		case 3:
//line lexer.rl:47
( lexer.act) = 5;
		case 4:
//line lexer.rl:51
( lexer.act) = 7;
		case 5:
//line lexer.rl:55
( lexer.act) = 9;
		case 6:
//line lexer.rl:59
( lexer.act) = 11;
		case 7:
//line lexer.rl:63
( lexer.act) = 13;
		case 8:
//line lexer.rl:67
( lexer.act) = 15;
		case 9:
//line lexer.rl:69
( lexer.act) = 16;
		case 10:
//line lexer.rl:73
( lexer.act) = 18;
		case 11:
//line lexer.rl:77
( lexer.act) = 20;
		case 12:
//line lexer.rl:81
( lexer.act) = 22;
		case 13:
//line lexer.rl:85
( lexer.act) = 24;
		case 14:
//line lexer.rl:95
( lexer.act) = 27;
		case 15:
//line lexer.rl:39
( lexer.te) = ( lexer.p)+1
{ token_kind = token_LPAREN; ( lexer.p)++; goto _out
 }
		case 16:
//line lexer.rl:41
( lexer.te) = ( lexer.p)+1
{ token_kind = token_RPAREN; ( lexer.p)++; goto _out
 }
		case 17:
//line lexer.rl:43
( lexer.te) = ( lexer.p)+1
{ token_kind = token_COMMA; ( lexer.p)++; goto _out
 }
		case 18:
//line lexer.rl:45
( lexer.te) = ( lexer.p)+1
{ token_kind = token_TEST_EQ; ( lexer.p)++; goto _out
 }
		case 19:
//line lexer.rl:49
( lexer.te) = ( lexer.p)+1
{ token_kind = token_TEST_NE; ( lexer.p)++; goto _out
 }
		case 20:
//line lexer.rl:57
( lexer.te) = ( lexer.p)+1
{ token_kind = token_TEST_GE; ( lexer.p)++; goto _out
 }
		case 21:
//line lexer.rl:65
( lexer.te) = ( lexer.p)+1
{ token_kind = token_TEST_LE; ( lexer.p)++; goto _out
 }
		case 22:
//line lexer.rl:71
( lexer.te) = ( lexer.p)+1
{ token_kind = token_TEST_MATCHES; ( lexer.p)++; goto _out
 }
		case 23:
//line lexer.rl:79
( lexer.te) = ( lexer.p)+1
{ token_kind = token_TEST_AND; ( lexer.p)++; goto _out
 }
		case 24:
//line lexer.rl:83
( lexer.te) = ( lexer.p)+1
{ token_kind = token_TEST_OR; ( lexer.p)++; goto _out
 }
		case 25:
//line lexer.rl:88
( lexer.te) = ( lexer.p)+1
{ /*lexer.ts++; lexer.te--;*/ token_kind = token_UNPARSED; ( lexer.p)++; goto _out
 }
		case 26:
//line lexer.rl:106
( lexer.te) = ( lexer.p)+1

		case 27:
//line lexer.rl:53
( lexer.te) = ( lexer.p)
( lexer.p)--
{ token_kind = token_TEST_GT; ( lexer.p)++; goto _out
 }
		case 28:
//line lexer.rl:61
( lexer.te) = ( lexer.p)
( lexer.p)--
{ token_kind = token_TEST_LT; ( lexer.p)++; goto _out
 }
		case 29:
//line lexer.rl:75
( lexer.te) = ( lexer.p)
( lexer.p)--
{ token_kind = token_TEST_NOT; ( lexer.p)++; goto _out
 }
		case 30:
//line lexer.rl:92
( lexer.te) = ( lexer.p)
( lexer.p)--
{ token_kind = token_UNPARSED; ( lexer.p)++; goto _out
 }
		case 31:
//line lexer.rl:95
( lexer.te) = ( lexer.p)
( lexer.p)--
{ if nameToId(string(lexer.data[lexer.ts : lexer.te])) != 0 {
			  		token_kind = token_FIELD
			  } else {
					token_kind = token_UNPARSED	
			  }; ( lexer.p)++; goto _out

			}
		case 32:
//line NONE:1
	switch ( lexer.act) {
	case 5:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_EQ; ( lexer.p)++; goto _out
 }
	case 7:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_NE; ( lexer.p)++; goto _out
 }
	case 9:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_GT; ( lexer.p)++; goto _out
 }
	case 11:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_GE; ( lexer.p)++; goto _out
 }
	case 13:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_LT; ( lexer.p)++; goto _out
 }
	case 15:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_LE; ( lexer.p)++; goto _out
 }
	case 16:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_CONTAINS; ( lexer.p)++; goto _out
 }
	case 18:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_MATCHES; ( lexer.p)++; goto _out
 }
	case 20:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_NOT; ( lexer.p)++; goto _out
 }
	case 22:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_AND; ( lexer.p)++; goto _out
 }
	case 24:
	{( lexer.p) = (( lexer.te)) - 1
 token_kind = token_TEST_OR; ( lexer.p)++; goto _out
 }
	case 27:
	{( lexer.p) = (( lexer.te)) - 1
 if nameToId(string(lexer.data[lexer.ts : lexer.te])) != 0 {
			  		token_kind = token_FIELD
			  } else {
					token_kind = token_UNPARSED	
			  }; ( lexer.p)++; goto _out

			}
	}
	
//line lexer.go:509
		}
	}

_again:
	_acts = int(_filterLexerImpl_to_state_actions[( lexer.cs)])
	_nacts = uint(_filterLexerImpl_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _filterLexerImpl_actions[_acts-1] {
		case 0:
//line NONE:1
( lexer.ts) = 0

//line lexer.go:523
		}
	}

	if ( lexer.cs) == 0 {
		goto _out
	}
	( lexer.p)++
	if ( lexer.p) != ( lexer.pe) {
		goto _resume
	}
	_test_eof: {}
	if ( lexer.p) == ( lexer.eof) {
		if _filterLexerImpl_eof_trans[( lexer.cs)] > 0 {
			_trans = int(_filterLexerImpl_eof_trans[( lexer.cs)] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//line lexer.rl:152
    if ( lexer.cs != filterLexerImpl_error ){
		lval.data = lexer.data[lexer.ts : lexer.te]
    }
	if filterDebug > 4 {
		fmt.Printf("Token text: %s\n", string(lval.data))
	}

	return token_kind
}

func (lexer *filterLexerImpl) Error(s string) {
	lexer.err = s
}