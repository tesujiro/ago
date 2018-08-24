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
	token		ast.Token
	rule		ast.Rule
	rules		[]ast.Rule
	pattern		ast.Pattern
	stmt		ast.Stmt
	opt_stmt	ast.Stmt
	stmts		[]ast.Stmt
	stmt_if		ast.Stmt
	expr		ast.Expr
	opt_expr	ast.Expr
	exprs		[]ast.Expr
	opt_exprs	[]ast.Expr
	ident_args	[]string
}

%type <rules>		program
%type <rule>		rule
%type <pattern> 	pattern
%type <stmt>		stmt
%type <opt_stmt>	opt_stmt
%type <stmts>		action
%type <stmts>		stmts
%type <stmt_if>		stmt_if
%type <expr>		expr
%type <opt_expr>	opt_expr
%type <exprs>		exprs
%type <opt_exprs>	opt_exprs
%type <ident_args>	ident_args

%token<token> IDENT NUMBER STRING TRUE FALSE NIL
%token<token> EQEQ NEQ GE LE ANDAND OROR LEN 
%token<token> PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ
%token<token> DELETE IN
%token<token> BEGIN END PRINT REGEXP
%token<token> IF ELSE FOR WHILE DO BREAK CONTINUE
%token<token> FUNC RETURN

%right '=' PLUSEQ MINUSEQ MULEQ DIVEQ  
%left OROR
%left ANDAND
%left IDENT
%left '~'
%left EQEQ NEQ
%left '>' '<' GE LE

%left '+' '-'
%left '*' '/' '%'
%right '!' UNARY
%left PLUSPLUS MINUSMINUS
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
	: pattern action
	{
		$$ = ast.Rule{Pattern: $1, Action: $2}
	}
	| pattern opt_semi opt_nls
	{
		$$ = ast.Rule{Pattern: $1, Action: []ast.Stmt{ &ast.PrintStmt{Exprs: defaultExprs }}}
	}
	| action
	{
		$$ = ast.Rule{Pattern: &ast.ExprPattern{}, Action: $1}
	}

pattern
	: FUNC IDENT '(' ident_args ')'
	{
		//fmt.Println("FUNC RULE")
		$$ = &ast.FuncPattern{Name: $2.Literal, Args: $4}
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
	| REGEXP ',' REGEXP
	{
		$$ = &ast.StartStopPattern{
			Start: &ast.MatchExpr{Expr: &defaultExpr, RegExpr: $1.Literal},
			Stop:  &ast.MatchExpr{Expr: &defaultExpr, RegExpr: $3.Literal},
		}
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
	| DELETE expr
	{
		$$ = &ast.DelStmt{Expr: $2}
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
	| FOR opt_stmt ';' opt_expr ';' opt_expr '{' stmts '}'
	{
		$$ = &ast.CForLoopStmt{Stmt1: $2, Expr2: $4, Expr3: $6, Stmts: $8}
	}
	| WHILE '{' stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $3}
	}
	| WHILE expr '{' stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $4, Expr: $2}
	}
	| DO '{' stmts '}' WHILE '(' expr ')'
	{
		$$ = &ast.DoLoopStmt{Stmts: $3, Expr: $7}
	}
	| BREAK
	{
		$$ = &ast.BreakStmt{}
	}
	| CONTINUE
	{
		$$ = &ast.ContinueStmt{}
	}
	| FOR '(' IDENT IN IDENT ')' '{' stmts '}'
	{
		$$ = &ast.HashLoopStmt{Key: $3.Literal, Hash: $5.Literal, Stmts:$8}
	}
	| RETURN opt_exprs
	{
		$$ = &ast.ReturnStmt{Exprs:$2}
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

opt_exprs
	: /* empty */
	{
		$$ = []ast.Expr{}
	}
	| exprs
	{
		$$ = $1
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
	| expr '[' exprs ']'
	{
		$$ = &ast.ItemExpr{Expr: $1, Index:$3}
	}
	/* REGEXP */
	| expr '~' REGEXP
	{
		$$ = &ast.MatchExpr{Expr: $1, RegExpr: $3.Literal}
	}
	| REGEXP
	{
		$$ = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: $1.Literal}
	}
	/* FUNCTION */
	| FUNC '(' ident_args ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $6}
	}
	| IDENT '(' opt_exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.Literal, SubExprs:$3}
	}
	| expr '(' opt_exprs ')'
	{
		$$ = &ast.AnonymousCallExpr{Expr: $1, SubExprs:$3}
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

ident_args
	: /* empty */
	{
		$$ = []string{}
	}
	| IDENT
	{
		$$ = []string{$1.Literal}
	}
	| ident_args ',' opt_nls IDENT
	{
		$$ = append($1,$4.Literal)
	}

nls
	: '\n'
	| nls '\n'

opt_stmt
	: /* empty */
	{
		$$ = nil
	}
	| stmt
	{
		$$ = $1
	}

opt_expr
	: /* empty */
	{
		$$ = nil
	}
	| expr
	{
		$$ = $1
	}

opt_nls
	: /* empty */
	| nls

semi
	: ';'

opt_semi
	: /* empty */
	| semi

%%
