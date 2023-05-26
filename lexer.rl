package gofilter

import "fmt"

%%{
	machine filterLexerImpl;
	ws      = ([\n\t\v\f ]);
	integer = "-"?([0-9]+ | [0-9]+'#'xdigit+);
	float = "-"? digit* '.' digit+;
	
	not_dquote_or_escape = [^"\\];
    escaped_something = /\\./;
    quote_string = '"' ( not_dquote_or_escape | escaped_something )* '"';
	
	DIGIT                       = "0".."9";
	HEXDIG                      = DIGIT | "A"i | "B"i | "C"i | "D"i | "E"i | "F"i;
	dec_octet                    = DIGIT | ( 0x31..0x39 DIGIT ) | ( "1" DIGIT{2} ) |
                                 ( "2" 0x30..0x34 DIGIT ) | ( "25" 0x30..0x35 );
  	IPv4address                  = dec_octet "." dec_octet "." dec_octet "." dec_octet;
	h16                          = HEXDIG{1,4};
  	ls32                         = ( h16 ":" h16 ) | IPv4address;
  	IPv6address                  = ( ( h16 ":" ){6} ls32 ) |
                                 ( "::" ( h16 ":" ){5} ls32 ) |
                                 ( h16? "::" ( h16 ":" ){4} ls32 ) |
                                 ( ( ( h16 ":" )? h16 )? "::" ( h16 ":" ){3} ls32 ) |
                                 ( ( ( h16 ":" ){,2} h16 )? "::" ( h16 ":" ){2} ls32 ) |
                                 ( ( ( h16 ":" ){,3} h16 )? "::" h16 ":" ls32 ) |
                                 ( ( ( h16 ":" ){,4} h16 )? "::" ls32 ) |
                                 ( ( ( h16 ":" ){,5} h16 )? "::" h16 ) |
                                 ( ( ( h16 ":" ){,6} h16 )? "::" );
	
	#field = ([0-9A-Za-z_.\-]*);
	
	#value = integer | float | quote_string | IPv4address | IPv6address;
	
	main :=
	|*
		"(" =>
			{ token_kind = token_LPAREN; fbreak; };
		")" =>
			{ token_kind = token_RPAREN; fbreak; };
		"," =>
			{ token_kind = token_COMMA; fbreak; };
		"==" =>
			{ token_kind = token_TEST_EQ; fbreak; };
		"eq" =>
			{ token_kind = token_TEST_EQ; fbreak; };
		"!=" =>
			{ token_kind = token_TEST_NE; fbreak; };
		"ne" =>
			{ token_kind = token_TEST_NE; fbreak; };
		">" =>
			{ token_kind = token_TEST_GT; fbreak; };
		"gt" =>
			{ token_kind = token_TEST_GT; fbreak; };
		">=" =>
			{ token_kind = token_TEST_GE; fbreak; };
		"ge" =>
			{ token_kind = token_TEST_GE; fbreak; };
		"<" =>
			{ token_kind = token_TEST_LT; fbreak; };
		"lt" =>
			{ token_kind = token_TEST_LT; fbreak; };
		"<=" =>
			{ token_kind = token_TEST_LE; fbreak; };
		"le" =>
			{ token_kind = token_TEST_LE; fbreak; };
		"contains" =>
			{ token_kind = token_TEST_CONTAINS; fbreak; };
		"~" =>
			{ token_kind = token_TEST_MATCHES; fbreak; };
		"matches" =>
			{ token_kind = token_TEST_MATCHES; fbreak; };
		"!" =>
			{ token_kind = token_TEST_NOT; fbreak; };
		"not" =>
			{ token_kind = token_TEST_NOT; fbreak; };
		"&&" =>
			{ token_kind = token_TEST_AND; fbreak; };
		"and" =>
			{ token_kind = token_TEST_AND; fbreak; };
		"||" =>
			{ token_kind = token_TEST_OR; fbreak; };
		"or" =>
			{ token_kind = token_TEST_OR; fbreak; };
			
		quote_string => 
			{ /*lexer.ts++; lexer.te--;*/ token_kind = token_UNPARSED; fbreak; };
		
		("-" | alnum | [_\.:])+ "/" digit+ =>
			# CIDR
			{ token_kind = token_UNPARSED; fbreak; };
		
		("-" | "+" | alnum | "_" | "." | ":")+
			{ if lexer.ctx.nameToId(string(lexer.data[lexer.ts : lexer.te])) != 0 {
			  		token_kind = token_FIELD
			  } else {
					token_kind = token_UNPARSED	
			  }; fbreak;
			};
		
		#value =>
		#	{ token_kind = token_UNPARSED; fbreak; };
		#field =>
		#	{ token_kind = token_FIELD; fbreak; };
		ws => {};
		
	*|;
}%%

%%{
	write data;
	variable data lexer.data;
	variable cs lexer.cs;
	variable p lexer.p;
	variable pe lexer.pe;
	variable act lexer.act;
	variable ts lexer.ts;
	variable te lexer.te;
	variable eof lexer.eof;
}%%


type filterLexerImpl struct {
	data []byte
	ctx  Context
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

func newLex(line []byte, ctx Context) *filterLexerImpl {
	lexer := filterLexerImpl{data: line, ctx: ctx}
	%%write init;
	lexer.pe = len(line)
	lexer.eof = len(line)
	return &lexer
}

//type FilterSymType struct {
//	Data []byte
//} // fake

func (lexer *filterLexerImpl) Lex(lval *filterSymType) int {
    token_kind := 0
	%% write exec;
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

func (lexer filterLexerImpl) Context() Context {
	return lexer.ctx
}
