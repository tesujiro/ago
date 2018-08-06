%{
	package parser
	import (
		//"fmt"
		"github.com/tesujiro/goa/ast"
	)

var defaultExpr = ast.FieldExpr{Expr: &ast.NumExpr{Literal: "0"}}
var defaultExprs = []ast.Expr{&defaultExpr}
//var IN_REGEXP bool
%}

%union{
	token ast.Token
	rule ast.Rule
	rules []ast.Rule
	pattern ast.Pattern
	stmt ast.Stmt
	stmts []ast.Stmt
	stmt_if ast.Stmt
	expr ast.Expr
	exprs []ast.Expr
}

%type <rules>	program
%type <rule>	rule
%type <pattern> pattern
%type <stmt>	stmt
%type <stmts>	action
%type <stmts>	stmts
%type <stmt_if>	stmt_if
%type <expr>	expr
%type <exprs>	exprs

%token<token> IDENT NUMBER STRING TRUE FALSE NIL
%token<token> EQEQ NEQ GE LE ANDAND OROR LEN 
%token<token> PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ
%token<token> BEGIN END PRINT REGEXP
%token<token> IF ELSE FOR BREAK CONTINUE
%token<token> FUNC RETURN

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
	: BEGIN action
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
	| BEGIN
	{
		$$ = &ast.BeginPattern{}
	}
	| END
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
	| PRINT 
	{
		$$ = &ast.PrintStmt{Exprs: defaultExprs }
	}
	| PRINT exprs
	{
		$$ = &ast.PrintStmt{Exprs: $2}
	}
	| stmt_if
	{
		$$ = $1
	}
	| FOR '{' stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $3}
	}
	| FOR expr '{' stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $4, Expr: $2}
	}
	| BREAK
	{
		$$ = &ast.BreakStmt{}
	}
	| CONTINUE
	{
		$$ = &ast.ContinueStmt{}
	}

/*
stmt_term
	: nls
	| semi opt_nls
*/

stmt_if
    : IF expr '{' stmts '}'
    {
        $$ = &ast.IfStmt{If: $2, Then: $4, Else: nil}
    }
    | stmt_if ELSE IF expr '{' stmts '}'
    {
            $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: $6} )
    }
    | stmt_if ELSE '{' stmts '}'
    {
        if $$.(*ast.IfStmt).Else != nil {
            yylex.Error("multiple else statement")
        } else {
            //$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
            $$.(*ast.IfStmt).Else = $4
        }
    }

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
	| TRUE
	{
		$$ = &ast.ConstExpr{Literal: $1.Literal}
	}
	| FALSE
	{
		$$ = &ast.ConstExpr{Literal: $1.Literal}
	}
	| NIL
	{
		$$ = &ast.ConstExpr{Literal: $1.Literal}
	}
	| '$' expr
	{
		$$ = &ast.FieldExpr{Expr: $2}
	}
	| STRING
	{
		$$ = &ast.StringExpr{Literal: $1.Literal}
	}
	| IDENT '[' exprs ']'
	{
		$$ = &ast.ItemExpr{Value: &ast.IdentExpr{Literal: $1.Literal}, Index:$3}
	}
	| expr '[' exprs ']'
	{
		$$ = &ast.ItemExpr{Value: $1, Index:$3}
	}
	/*
	| '/' 
	{
		fmt.Println("path1")
		IN_REGEXP=true
	}
	| REGEXP
	{
		fmt.Println("path2:",$1.Literal)
		IN_REGEXP=false
		$$ = &ast.IdentExpr{Literal: $1.Literal}
	}
	*/
	/* UNARY EXPRESSION */
	| '+' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "+", Expr:$2}
	}
	| '-' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "-", Expr:$2}
	}
	| '!' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "!", Expr:$2}
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
	/* BOOL EXPRESSION */
	| expr OROR expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "||", Right: $3}
	}
	| expr ANDAND expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "&&", Right: $3}
	}
	// TODO: 'NOT'
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
