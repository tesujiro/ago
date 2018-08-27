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
	stmts		[]ast.Stmt
	expr		ast.Expr
	exprs		[]ast.Expr
	ident_args	[]string
}

%type <rules>		program
%type <rule>		rule
%type <pattern> 	pattern
%type <stmt>		stmt
%type <stmt>		opt_stmt
%type <stmts>		action
%type <stmts>		stmts
%type <stmts>		opt_stmts
%type <stmt>		stmt_if
%type <expr>		expr
%type <expr>		simp_expr
%type <expr>		variable
%type <expr>		opt_expr
%type <exprs>		exprs
%type <exprs>		opt_exprs
%type <ident_args>	ident_args

%token <token> IDENT NUMBER STRING TRUE FALSE NIL
%token <token> EQEQ NEQ GE LE ANDAND OROR LEN 
%token <token> PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ
%token <token> DELETE IN
%token <token> BEGIN END PRINT REGEXP
%token <token> IF ELSE FOR WHILE DO BREAK CONTINUE
%token <token> FUNC RETURN

%right '=' PLUSEQ MINUSEQ MULEQ DIVEQ  
%left OROR
%left ANDAND
%left IDENT
%left '~'
%left EQEQ NEQ
%left '>' '<' GE LE

%left CONCAT_OP
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
	: pattern action opt_term
	{
		$$ = ast.Rule{Pattern: $1, Action: $2}
	}
	| pattern term
	{
		$$ = ast.Rule{Pattern: $1, Action: []ast.Stmt{ &ast.PrintStmt{Exprs: defaultExprs }}}
	}
	/*
	| action
	{
		$$ = ast.Rule{Pattern: &ast.ExprPattern{}, Action: $1}
	}
	*/

pattern
	: /* empty */
	{
		$$ = &ast.ExprPattern{}
	}
	| FUNC IDENT '(' ident_args ')'
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
	: '{' opt_stmts '}'
	{
		$$ = $2
	}

opt_stmts
	: /* empty */
	{
		$$ = []ast.Stmt{}
	}
	| stmts opt_term
	{
		$$ = $1
	}
	
stmts
	: opt_term stmt
	{
		$$ = []ast.Stmt{$2}
	}
	| stmts semi opt_nls stmt
	{
		$$ = append($1,$4)
	}

stmt
	: variable '=' expr
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
	| FOR '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $3}
	}
	| FOR expr '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $4, Expr: $2}
	}
	| FOR opt_stmt ';' opt_expr ';' opt_expr '{' opt_stmts '}'
	{
		$$ = &ast.CForLoopStmt{Stmt1: $2, Expr2: $4, Expr3: $6, Stmts: $8}
	}
	| WHILE '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $3}
	}
	| WHILE expr '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $4, Expr: $2}
	}
	| DO '{' opt_stmts '}' WHILE '(' expr ')'
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
	| FOR '(' IDENT IN IDENT ')' '{' opt_stmts '}'
	{
		$$ = &ast.HashLoopStmt{Key: $3.Literal, Hash: $5.Literal, Stmts:$8}
	}
	| RETURN opt_exprs
	{
		$$ = &ast.ReturnStmt{Exprs:$2}
	}

stmt_if
	: IF expr '{' opt_stmts '}'
	{
	    $$ = &ast.IfStmt{If: $2, Then: $4, Else: nil}
	}
	| stmt_if ELSE IF expr '{' opt_stmts '}'
	{
	        $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: $6} )
	}
	| stmt_if ELSE '{' opt_stmts '}'
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
	: simp_expr
	{
		$$ = $1
	}
	/*
	| expr simp_expr %prec CONCAT_OP
	{
		$$ = &ast.ConcatExpr{Left: $1, Right: $2}
	}
	*/
	/* FUNCTION DEFINITION */
	| FUNC '(' ident_args ')' '{' opt_stmts '}'
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $6}
	}
	/* COMPOSITE EXPRESSION */
	| variable PLUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "+=", Right: $3}
	}
	| variable MINUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "-=", Right: $3}
	}
	| variable MULEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "*=", Right: $3}
	}
	| variable DIVEQ expr
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


simp_expr
	: variable
	{
		$$ = $1
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
	| STRING
	{
		$$ = &ast.StringExpr{Literal: $1.Literal}
	}
	/* REGEXP */
	| simp_expr '~' REGEXP
	{
		$$ = &ast.MatchExpr{Expr: $1, RegExpr: $3.Literal}
	}
	| REGEXP
	{
		$$ = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: $1.Literal}
	}
	/* FUNCTION CALL */
	| IDENT '(' opt_exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.Literal, SubExprs:$3}
	}
	| expr '(' opt_exprs ')'
	{
		$$ = &ast.AnonymousCallExpr{Expr: $1, SubExprs:$3}
	}
	/* COMPOSITE EXPRESSION */
	| PLUSPLUS expr
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "++"}
	}
	| expr PLUSPLUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "++", After:true}
	}
	| MINUSMINUS expr
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "--"}
	}
	| expr MINUSMINUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "--", After:true}
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

variable
	: IDENT
	{
		$$ = &ast.IdentExpr{Literal: $1.Literal}
	}
	| '$' expr
	{
		$$ = &ast.FieldExpr{Expr: $2}
	}
	| simp_expr '[' exprs ']'
	{
		$$ = &ast.ItemExpr{Expr: $1, Index:$3}
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

opt_term
	: /* empty */
	| term

term
	: semi nls
	| nls
	| semi

semi
	: ';'

opt_nls
	: /* empty */
	| nls

nls
	: '\n'
	| nls '\n'

%%
