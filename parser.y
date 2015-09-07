%{

package gofilter

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"strconv"
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
			return  str, true
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

%}

%union {
	nnode node
	data []byte
}

%token token_TEST_NOT token_TEST_AND token_TEST_OR
%token token_LPAREN token_RPAREN token_COMMA
%token token_TEST_EQ token_TEST_NE 			
%token token_TEST_GT token_TEST_GE token_TEST_LT token_TEST_LE
%token token_TEST_CONTAINS token_TEST_MATCHES

%token token_FIELD token_STRING token_UNPARSED


%left token_TEST_AND
%left token_TEST_OR
%left token_TEST_NOT
%left token_TEST_EQ token_TEST_NE


%%
search_codition:
		search_codition token_TEST_AND search_codition
		{ 
			result := nodeAnd{left: $1.nnode, right: $3.nnode}
			$$.nnode = &result
			setResult(filterlex, &result)
		}
	
	|	search_codition token_TEST_OR search_codition
		{ 
			result := nodeOr{left: $1.nnode, right: $3.nnode}
			$$.nnode = &result
			setResult(filterlex, &result)
		}
	
	| 	token_TEST_NOT search_codition
		{
			result := nodeNot{right: $2.nnode}
			$$.nnode = &result
			setResult(filterlex, &result)
		}
	
	| 	token_LPAREN search_codition token_RPAREN
	 	{ 
			$$.nnode = $2.nnode
			setResult(filterlex, $2.nnode)
		}	
	| 	predicate
	;
		
predicate:
	token_FIELD token_TEST_EQ token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_EQ, $3.data); ok {
				$$.nnode = &nodeEq{fieldName: (string)($1.data), value: val}
				setResult(filterlex, $$.nnode)	
			} else {
				return 1
			}
		}
	
	| 	token_FIELD token_TEST_NE token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_NE, $3.data); ok {
				$$.nnode = &nodeNot{right: &nodeEq{fieldName: (string)($1.data), value: val}}
				setResult(filterlex, $$.nnode)
			} else {
				return 1
			}
		}
		
	| 	token_FIELD token_TEST_GT token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_GT, $3.data); ok {
				$$.nnode = &nodeGt{fieldName: (string)($1.data), value: val}
				setResult(filterlex, $$.nnode)
			} else {
				return 1
			}
		}
		
	| 	token_FIELD token_TEST_GE token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_GE, $3.data); ok {
				$$.nnode = &nodeGe{fieldName: (string)($1.data), value: val}
				setResult(filterlex, $$.nnode)
			} else {
				return 1
			}
		}
		
	| 	token_FIELD token_TEST_LT token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_LT, $3.data); ok {
				$$.nnode = &nodeLt{fieldName: (string)($1.data), value: val}
				setResult(filterlex, $$.nnode)
			} else {
				return 1
			}
		}
		
	| 	token_FIELD token_TEST_LE token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_LE, $3.data); ok {
				$$.nnode = &nodeLe{fieldName: (string)($1.data), value: val}
				setResult(filterlex, $$.nnode)
			} else {
				return 1
			}
		}
		
	| 	token_FIELD token_TEST_CONTAINS token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_CONTAINS, $3.data); ok {
				$$.nnode = &nodeContains{fieldName: (string)($1.data), value: val}
				setResult(filterlex, $$.nnode)
			} else {
				return 1
			}
		}
		
	| 	token_FIELD token_TEST_MATCHES token_UNPARSED
		{
			if val, ok := checkFieldNameVsTypeValue(filterlex, (string)($1.data), token_TEST_MATCHES, $3.data); ok {
				r_expr, err := regexp.Compile(val.(string))
				if err != nil {
					str := fmt.Sprintf("Incorrect reqular expresstion \"%s\": %s.", (string)($3.data), err)
					filterlex.Error(str)
					return 1
				}
				
				$$.nnode = &nodeMatch{fieldName: (string)($1.data), reg_expr: r_expr}
				setResult(filterlex, $$.nnode)
			} else {
				return 1
			}
		}
	
	|	token_FIELD
		{
			$$.nnode = &nodeExist{fieldName: (string)($1.data)}
			setResult(filterlex, $$.nnode)
		}
	
	|	token_UNPARSED
		{
			str := fmt.Sprintf("Field with name \"%s\" does not exists.", (string)($1.data))
			filterlex.Error(str)
			return 1

		}
	;

%%


