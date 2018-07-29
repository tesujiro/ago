%{
	package parser
	import (
		//"fmt"
		"github.com/tesujiro/goa/ast"
	)

var defaultExpr = ast.FieldExpr{Expr: &ast.NumExpr{Literal: "0"}}
var defaultExprs = []ast.Expr{&defaultExpr}
%}

%union{
	token ast.Token
	rule ast.Rule
	rules []ast.Rule
	pattern ast.Pattern
	stmt ast.Stmt
	stmts []ast.Stmt
	expr ast.Expr
	exprs []ast.Expr
}

%type <rules>	program
%type <rule>	rule
%type <pattern> pattern
%type <stmts>	action
%type <stmts>	stmts
%type <stmt>	stmt
%type <expr>	expr
%type <exprs>	exprs

%token<token> IDENT NUMBER STRING BEGIN END LEX_BEGIN LEX_END LEX_PRINT TRUE FALSE NIL FUNC RETURN EQEQ NEQ GE LE IF ELSE ANDAND OROR LEN FOR BREAK CONTINUE PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ

%right '='
%left OROR
%left ANDAND
%left IDENT
%left EQEQ NEQ
%left '>' '<' GE LE

%left '+' '-' PLUSPLUS MINUSMINUS
%left '*' '/' '%'
%right UNARY
%left '$'
%left '(' ')'

%%

program
	: /* empty */
	{
		$$ = []ast.Rule{}
	}
	| program rule
	{
		$$ = append($1,$2)
		yylex.(*Lexer).result = $$
	}

rule
	/*
	: LEX_BEGIN action
	{
		$$ = &ast.BeginRule{Pattern: $1, Action: $2}
	}
	*/
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
		$$ = &ast.ExprPattern{}
	}
	| LEX_BEGIN
	{
		$$ = &ast.BeginPattern{}
	}
	| LEX_END
	{
		$$ = &ast.EndPattern{}
	}
	| expr
	{
		$$ = &ast.ExprPattern{Expr:$1}
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
	| stmts stmt opt_semi opt_nls
	{
		$$ = append($1,$2)
	}

stmt
	: expr '=' expr
	{
		$$ = &ast.AssStmt{Left: []ast.Expr{$1}, Right: []ast.Expr{$3}}
	}
	| exprs '=' exprs
	{
		$$ = &ast.AssStmt{Left: $1, Right: $3}
	}
	| expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| LEX_PRINT 
	{
		$$ = &ast.PrintStmt{Exprs: defaultExprs }
	}
	| LEX_PRINT exprs
	{
		$$ = &ast.PrintStmt{Exprs: $2}
	}

/*
stmt_term
	: nls
	| semi opt_nls
*/

exprs
	: expr
	{
		$$ = []ast.Expr{$1}
	}
	| exprs ',' opt_nls expr
	{
		$$ = append($1,$4)
	}

expr
	: IDENT
	{
		$$ = &ast.IdentExpr{Literal: $1.Literal}
	}
	| NUMBER
	{
		$$ = &ast.NumExpr{Literal: $1.Literal}
	}
	| '$' expr
	{
		$$ = &ast.FieldExpr{Expr: $2}
	}
	| STRING
	{
		$$ = &ast.StringExpr{Literal: $1.Literal}
	}
	/* COMPOSITE EXPRESSION */
	| expr PLUSPLUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "++"}
	}
	| expr MINUSMINUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "--"}
	}
	| expr PLUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "+=", Right: $3}
	}
	| expr MINUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "-=", Right: $3}
	}
	| expr MULEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "*=", Right: $3}
	}
	| expr DIVEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "/=", Right: $3}
	}
	/* RELATION EXPRESSION */
	| expr EQEQ expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "==", Right: $3}
	}
	| expr NEQ expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "!=", Right: $3}
	}
	| expr '>' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">", Right: $3}
	}
	| expr GE expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">=", Right: $3}
	}
	| expr '<' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<", Right: $3}
	}
	| expr LE expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<=", Right: $3}
	}
	/* ARITHMETIC EXPRESSION */
	| '(' expr ')'
	{
		$$ = &ast.ParentExpr{SubExpr: $2}
	}
	| expr '+' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "+", Right: $3}
	}
	| expr '-' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "-", Right: $3}
	}
	| expr '*' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "*", Right: $3}
	}
	| expr '/' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "/", Right: $3}
	}
	| expr '%' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "%", Right: $3}
	}

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
