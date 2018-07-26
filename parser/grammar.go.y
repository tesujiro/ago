%{
	package parser
	import (
		//"fmt"
		"github.com/tesujiro/goa/ast"
	)
%}

%union{
	token ast.Token
	rule ast.Rule
	rules []ast.Rule
	pattern ast.Pattern
	stmt ast.Stmt
	stmts []ast.Stmt
}

%type <rules>	program
%type <rule>	rule
%type <pattern> pattern
%type <stmts>	action
%type <stmts>	stmts
%type <stmt>	stmt

%token<token> IDENT NUMBER STRING BEGIN END LEX_BEGIN LEX_END LEX_PRINT

%%

program
	: /* empty */
	{
		$$ = []ast.Rule{}
	}
	/*
	| rule 
	{
		$$ = []ast.Rule{$1}
		fmt.Println("PROGRAM EMPTY:",$$)
	}
	*/
	| program rule
	{
		$$ = append($1,$2)
		yylex.(*Lexer).result = $$
	}

rule
	: pattern action
	{
		$$ = ast.Rule{Pattern: $1, Action: $2}
	}
	/*
	| pattern stmt_term
	*/

pattern
	: /* empty */
	{
		$$ = ast.Pattern{}
	}
	| LEX_BEGIN
	{
		$$ = ast.Pattern{}
	}
	| LEX_END
	{
		$$ = ast.Pattern{}
	}
	| '$'
	{
		$$ = ast.Pattern{}
	}

action
	: '{' stmts '}' opt_semi opt_nls
	{
		$$ = $2
	}

stmts
	: /* empty */
	{
		$$ = []ast.Stmt{}
	}
	/*
	| stmt opt_semi opt_nls
	{
		$$ = []ast.Stmt{$1}
	}
	*/
	| stmts stmt opt_semi opt_nls
	{
		$$ = append($1,$2)
	}

stmt
	: LEX_PRINT
	{
		$$ = ast.Stmt{Message:"print"}
	}

/*
stmt_term
	: nls
	| semi opt_nls
*/

nls
	: '\n'
	| nls '\n'

opt_nls
	: /* empty */
	| nls

semi
	: ';'

opt_semi
	: /* empty */
	| semi

%%

