//line parser.y:1

package gofilter

import __yyfmt__ "fmt"

//line parser.y:4
import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

func setResult(filterlex filterLexer, n node) {
	f := filterlex.(*filterLexerImpl)
	f.result = n
}

func makeParseNumErrStr(err error) string {
	e := err.(*strconv.NumError)
	return e.Err.Error() + " " + strconv.Quote(e.Num)
}

func checkFieldNameVsTypeValue(filterlex filterLexer, name string, op int, data []byte) (interface{}, bool) {
	data_str := string(data)
	field_type := nameToFieldType(name)

	switch {
	// bool
	case field_type == FT_BOOL:
		val, err := strconv.ParseBool(data_str)
		if err == nil {
			return val, true
		}
		filterlex.Error(makeParseNumErrStr(err))

	// uint
	case field_type == FT_UINT:
		val, err := strconv.ParseUint(data_str, 0, 0)
		if err == nil {
			return uint(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_UINT8:
		val, err := strconv.ParseUint(data_str, 0, 8)
		if err == nil {
			return uint8(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_UINT16:
		val, err := strconv.ParseUint(data_str, 0, 16)
		if err == nil {
			return uint16(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_UINT32:
		val, err := strconv.ParseUint(data_str, 0, 32)
		if err == nil {
			return uint32(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_UINT64:
		val, err := strconv.ParseUint(data_str, 0, 64)
		if err == nil {
			return uint64(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))

	// int
	case field_type == FT_INT:
		val, err := strconv.ParseInt(data_str, 0, 0)
		if err == nil {
			return int(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_INT8:
		val, err := strconv.ParseInt(data_str, 0, 8)
		if err == nil {
			return int8(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_INT16:
		val, err := strconv.ParseInt(data_str, 0, 16)
		if err == nil {
			return int16(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_INT32:
		val, err := strconv.ParseInt(data_str, 0, 32)
		if err == nil {
			return int32(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_INT64:
		val, err := strconv.ParseInt(data_str, 0, 64)
		if err == nil {
			return int64(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))

	// float
	case field_type == FT_FLOAT32:
		val, err := strconv.ParseFloat(data_str, 32)
		if err == nil {
			return float32(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))
	case field_type == FT_FLOAT64:
		val, err := strconv.ParseFloat(data_str, 64)
		if err == nil {
			return float64(val), true
		}
		filterlex.Error(makeParseNumErrStr(err))

	// string or []byte
	case field_type == FT_STRING || field_type == FT_BYTES:
		if str, err := strconv.Unquote(data_str); err == nil {
			return str, true
		}
		if b, ok := parseBytes(data_str); ok {
			return b, true
		}
		return data_str, true

	// ip
	case field_type == FT_IP:
		ip := net.ParseIP(data_str)
		if ip != nil {
			return ip.To16(), true
		} else if _, ipnet, err := net.ParseCIDR(data_str); err == nil {
			ipnet.IP = ipnet.IP.To16()
			return ipnet, true
		}
		filterlex.Error("invalid IP address " + data_str)

	// mac
	case field_type == FT_MAC:
		mac, err := net.ParseMAC(data_str)
		if err == nil {
			return mac, true
		} else if b, ok := parseBytes(data_str); ok && op == token_TEST_CONTAINS {
			return b, true
		}
		filterlex.Error(err.Error())

	default:
		filterlex.Error("can not compare field " + strconv.Quote(name) +
			" and " + strconv.Quote(data_str))
	}

	return nil, false
}

func parseBytes(s string) (b []byte, ok bool) {
	l := strings.Split(s, ":")
	b = make([]byte, 0)
	for _, x := range l {
		if i, err := strconv.ParseUint(x, 16, 8); err == nil {
			b = append(b, byte(i))
		} else {
			return []byte{}, false
		}
	}

	return b, true
}

//line parser.y:169
type filterSymType struct {
	yys   int
	nnode node
	data  []byte
}

const token_TEST_NOT = 57346
const token_TEST_AND = 57347
const token_TEST_OR = 57348
const token_LPAREN = 57349
const token_RPAREN = 57350
const token_COMMA = 57351
const token_TEST_EQ = 57352
const token_TEST_NE = 57353
const token_TEST_GT = 57354
const token_TEST_GE = 57355
const token_TEST_LT = 57356
const token_TEST_LE = 57357
const token_TEST_CONTAINS = 57358
const token_TEST_MATCHES = 57359
const token_FIELD = 57360
const token_STRING = 57361
const token_UNPARSED = 57362

var filterToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_TEST_NOT",
	"token_TEST_AND",
	"token_TEST_OR",
	"token_LPAREN",
	"token_RPAREN",
	"token_COMMA",
	"token_TEST_EQ",
	"token_TEST_NE",
	"token_TEST_GT",
	"token_TEST_GE",
	"token_TEST_LT",
	"token_TEST_LE",
	"token_TEST_CONTAINS",
	"token_TEST_MATCHES",
	"token_FIELD",
	"token_STRING",
	"token_UNPARSED",
}
var filterStatenames = [...]string{}

const filterEofCode = 1
const filterErrCode = 2
const filterMaxDepth = 200

//line parser.y:323

//line yacctab:1
var filterExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const filterNprod = 16
const filterPrivate = 57344

var filterTokenNames []string
var filterStates []string

const filterLast = 37

var filterAct = [...]int{

	11, 12, 13, 14, 15, 16, 17, 18, 2, 29,
	4, 3, 28, 27, 26, 25, 24, 23, 22, 7,
	8, 8, 5, 1, 6, 0, 9, 10, 0, 0,
	0, 19, 20, 7, 8, 0, 21,
}
var filterPact = [...]int{

	4, 14, 4, 4, -1000, -10, -1000, 4, 4, -1000,
	28, -2, -3, -4, -5, -6, -7, -8, -11, 15,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var filterPgo = [...]int{

	0, 23, 10,
}
var filterR1 = [...]int{

	0, 1, 1, 1, 1, 1, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2,
}
var filterR2 = [...]int{

	0, 3, 3, 2, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 1, 1,
}
var filterChk = [...]int{

	-1000, -1, 4, 7, -2, 18, 20, 5, 6, -1,
	-1, 10, 11, 12, 13, 14, 15, 16, 17, -1,
	-1, 8, 20, 20, 20, 20, 20, 20, 20, 20,
}
var filterDef = [...]int{

	0, -2, 0, 0, 5, 14, 15, 0, 0, 3,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	2, 4, 6, 7, 8, 9, 10, 11, 12, 13,
}
var filterTok1 = [...]int{

	1,
}
var filterTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}
var filterTok3 = [...]int{
	0,
}

var filterErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	filterDebug        = 0
	filterErrorVerbose = false
)

type filterLexer interface {
	Lex(lval *filterSymType) int
	Error(s string)
}

type filterParser interface {
	Parse(filterLexer) int
	Lookahead() int
}

type filterParserImpl struct {
	lookahead func() int
}

func (p *filterParserImpl) Lookahead() int {
	return p.lookahead()
}

func filterNewParser() filterParser {
	p := &filterParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const filterFlag = -1000

func filterTokname(c int) string {
	if c >= 1 && c-1 < len(filterToknames) {
		if filterToknames[c-1] != "" {
			return filterToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func filterStatname(s int) string {
	if s >= 0 && s < len(filterStatenames) {
		if filterStatenames[s] != "" {
			return filterStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func filterErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !filterErrorVerbose {
		return "syntax error"
	}

	for _, e := range filterErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + filterTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := filterPact[state]
	for tok := TOKSTART; tok-1 < len(filterToknames); tok++ {
		if n := base + tok; n >= 0 && n < filterLast && filterChk[filterAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if filterDef[state] == -2 {
		i := 0
		for filterExca[i] != -1 || filterExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; filterExca[i] >= 0; i += 2 {
			tok := filterExca[i]
			if tok < TOKSTART || filterExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if filterExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += filterTokname(tok)
	}
	return res
}

func filterlex1(lex filterLexer, lval *filterSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = filterTok1[0]
		goto out
	}
	if char < len(filterTok1) {
		token = filterTok1[char]
		goto out
	}
	if char >= filterPrivate {
		if char < filterPrivate+len(filterTok2) {
			token = filterTok2[char-filterPrivate]
			goto out
		}
	}
	for i := 0; i < len(filterTok3); i += 2 {
		token = filterTok3[i+0]
		if token == char {
			token = filterTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = filterTok2[1] /* unknown char */
	}
	if filterDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", filterTokname(token), uint(char))
	}
	return char, token
}

func filterParse(filterlex filterLexer) int {
	return filterNewParser().Parse(filterlex)
}

func (filterrcvr *filterParserImpl) Parse(filterlex filterLexer) int {
	var filtern int
	var filterlval filterSymType
	var filterVAL filterSymType
	var filterDollar []filterSymType
	_ = filterDollar // silence set and not used
	filterS := make([]filterSymType, filterMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	filterstate := 0
	filterchar := -1
	filtertoken := -1 // filterchar translated into internal numbering
	filterrcvr.lookahead = func() int { return filterchar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		filterstate = -1
		filterchar = -1
		filtertoken = -1
	}()
	filterp := -1
	goto filterstack

ret0:
	return 0

ret1:
	return 1

filterstack:
	/* put a state and value onto the stack */
	if filterDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", filterTokname(filtertoken), filterStatname(filterstate))
	}

	filterp++
	if filterp >= len(filterS) {
		nyys := make([]filterSymType, len(filterS)*2)
		copy(nyys, filterS)
		filterS = nyys
	}
	filterS[filterp] = filterVAL
	filterS[filterp].yys = filterstate

filternewstate:
	filtern = filterPact[filterstate]
	if filtern <= filterFlag {
		goto filterdefault /* simple state */
	}
	if filterchar < 0 {
		filterchar, filtertoken = filterlex1(filterlex, &filterlval)
	}
	filtern += filtertoken
	if filtern < 0 || filtern >= filterLast {
		goto filterdefault
	}
	filtern = filterAct[filtern]
	if filterChk[filtern] == filtertoken { /* valid shift */
		filterchar = -1
		filtertoken = -1
		filterVAL = filterlval
		filterstate = filtern
		if Errflag > 0 {
			Errflag--
		}
		goto filterstack
	}

filterdefault:
	/* default state action */
	filtern = filterDef[filterstate]
	if filtern == -2 {
		if filterchar < 0 {
			filterchar, filtertoken = filterlex1(filterlex, &filterlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if filterExca[xi+0] == -1 && filterExca[xi+1] == filterstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			filtern = filterExca[xi+0]
			if filtern < 0 || filtern == filtertoken {
				break
			}
		}
		filtern = filterExca[xi+1]
		if filtern < 0 {
			goto ret0
		}
	}
	if filtern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			filterlex.Error(filterErrorMessage(filterstate, filtertoken))
			Nerrs++
			if filterDebug >= 1 {
				__yyfmt__.Printf("%s", filterStatname(filterstate))
				__yyfmt__.Printf(" saw %s\n", filterTokname(filtertoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for filterp >= 0 {
				filtern = filterPact[filterS[filterp].yys] + filterErrCode
				if filtern >= 0 && filtern < filterLast {
					filterstate = filterAct[filtern] /* simulate a shift of "error" */
					if filterChk[filterstate] == filterErrCode {
						goto filterstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if filterDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", filterS[filterp].yys)
				}
				filterp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if filterDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", filterTokname(filtertoken))
			}
			if filtertoken == filterEofCode {
				goto ret1
			}
			filterchar = -1
			filtertoken = -1
			goto filternewstate /* try again in the same state */
		}
	}

	/* reduction by production filtern */
	if filterDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", filtern, filterStatname(filterstate))
	}

	filternt := filtern
	filterpt := filterp
	_ = filterpt // guard against "declared and not used"

	filterp -= filterR2[filtern]
	// filterp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if filterp+1 >= len(filterS) {
		nyys := make([]filterSymType, len(filterS)*2)
		copy(nyys, filterS)
		filterS = nyys
	}
	filterVAL = filterS[filterp+1]

	/* consult goto table to find next state */
	filtern = filterR1[filtern]
	filterg := filterPgo[filtern]
	filterj := filterg + filterS[filterp].yys + 1

	if filterj >= filterLast {
		filterstate = filterAct[filterg]
	} else {
		filterstate = filterAct[filterj]
		if filterChk[filterstate] != -filtern {
			filterstate = filterAct[filterg]
		}
	}
	// dummy call; replaced with literal code
	switch filternt {

	case 1:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:192
		{
			result := nodeAnd{left: filterDollar[1].nnode, right: filterDollar[3].nnode}
			filterVAL.nnode = &result
			setResult(filterlex, &result)
		}
	case 2:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:199
		{
			result := nodeOr{left: filterDollar[1].nnode, right: filterDollar[3].nnode}
			filterVAL.nnode = &result
			setResult(filterlex, &result)
		}
	case 3:
		filterDollar = filterS[filterpt-2 : filterpt+1]
		//line parser.y:206
		{
			result := nodeNot{right: filterDollar[2].nnode}
			filterVAL.nnode = &result
			setResult(filterlex, &result)
		}
	case 4:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:213
		{
			filterVAL.nnode = filterDollar[2].nnode
			setResult(filterlex, filterDollar[2].nnode)
		}
	case 6:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:222
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_EQ, filterDollar[3].data); ok {
				filterVAL.nnode = &nodeEq{fieldName: (string)(filterDollar[1].data), value: val}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 7:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:232
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_NE, filterDollar[3].data); ok {
				filterVAL.nnode = &nodeNot{right: &nodeEq{fieldName: (string)(filterDollar[1].data), value: val}}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 8:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:242
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_GT, filterDollar[3].data); ok {
				filterVAL.nnode = &nodeGt{fieldName: (string)(filterDollar[1].data), value: val}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 9:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:252
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_GE, filterDollar[3].data); ok {
				filterVAL.nnode = &nodeGe{fieldName: (string)(filterDollar[1].data), value: val}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 10:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:262
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_LT, filterDollar[3].data); ok {
				filterVAL.nnode = &nodeLt{fieldName: (string)(filterDollar[1].data), value: val}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 11:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:272
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_LE, filterDollar[3].data); ok {
				filterVAL.nnode = &nodeLe{fieldName: (string)(filterDollar[1].data), value: val}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 12:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:282
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_CONTAINS, filterDollar[3].data); ok {
				filterVAL.nnode = &nodeContains{fieldName: (string)(filterDollar[1].data), value: val}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 13:
		filterDollar = filterS[filterpt-3 : filterpt+1]
		//line parser.y:292
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)(filterDollar[1].data), token_TEST_MATCHES, filterDollar[3].data); ok {
				r_expr, err := regexp.Compile(val.(string))
				if err != nil {
					str := fmt.Sprintf("Incorrect reqular expresstion \"%s\": %s.", (string)(filterDollar[3].data), err)
					filterlex.Error(str)
					return 1
				}

				filterVAL.nnode = &nodeMatch{fieldName: (string)(filterDollar[1].data), reg_expr: r_expr}
				setResult(filterlex, filterVAL.nnode)
			} else {
				return 1
			}
		}
	case 14:
		filterDollar = filterS[filterpt-1 : filterpt+1]
		//line parser.y:309
		{
			filterVAL.nnode = &nodeExist{fieldName: (string)(filterDollar[1].data)}
			setResult(filterlex, filterVAL.nnode)
		}
	case 15:
		filterDollar = filterS[filterpt-1 : filterpt+1]
		//line parser.y:315
		{
			str := fmt.Sprintf("Field with name \"%s\" does not exists.", (string)(filterDollar[1].data))
			filterlex.Error(str)
			return 1

		}
	}
	goto filterstack /* stack new state and value */
}
