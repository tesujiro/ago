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
%type <expr>		common_expr
%type <expr>		multi_val_expr
%type <expr>		simp_expr
%type <expr>		non_post_simp_expr
%type <expr>		variable
%type <expr>		opt_expr
%type <exprs>		exprs
%type <exprs>		variables
%type <exprs>		opt_exprs
%type <ident_args>	ident_args

%token <token> IDENT NUMBER STRING TRUE FALSE NIL
%token <token> EQEQ NEQ GE LE ANDAND OROR LEN 
%token <token> PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ
%token <token> DELETE IN
%token <token> BEGIN END PRINT REGEXP
%token <token> IF ELSE FOR WHILE DO BREAK CONTINUE
%token <token> FUNC RETURN EXIT NEXT
%token <token> CONCAT_OP

%right '=' PLUSEQ MINUSEQ MULEQ DIVEQ
%right '?' ':'
%left OROR
%left ANDAND
/*%left IDENT*/
%nonassoc ',' vars
%left '~'
%left EQEQ NEQ
%left '>' '<' GE LE

%left CONCAT_OP
%left STRING NUMBER
%left '+' '-'
%left '*' '/' '%'
%right '!' UNARY
%left PLUSPLUS MINUSMINUS
%left '$'
%nonassoc '['
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
	| action opt_term
	{
		$$ = ast.Rule{Pattern: &ast.ExprPattern{}, Action: $1}
	}

pattern
	/*
	:
	{
		$$ = &ast.ExprPattern{}
	}
	*/
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
	| stmts semi opt_term stmt
	{
		$$ = append($1,$4)
	}

stmt
	: expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| multi_val_expr
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
	| NEXT
	{
		$$ = &ast.NextStmt{}
	}
	| FOR '(' IDENT IN IDENT ')' '{' opt_stmts '}'
	{
		$$ = &ast.HashLoopStmt{Key: $3.Literal, Hash: $5.Literal, Stmts:$8}
	}
	| RETURN opt_exprs
	{
		$$ = &ast.ReturnStmt{Exprs:$2}
	}
	| EXIT opt_expr
	{
		$$ = &ast.ExitStmt{Expr:$2}
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
	/*
	: expr
	: simp_expr
	: non_post_simp_expr
	*/
	: expr
	{
		$$ = []ast.Expr{$1}
	}
	/*
	| exprs ',' opt_term expr
	| exprs ',' opt_term simp_expr
	| exprs ',' opt_term non_post_simp_expr
	*/
	| exprs ',' opt_term expr
	{
		$$ = append($1,$4)
	}

multi_val_expr
	: variables '=' exprs
	{
		$$ = &ast.AssExpr{Left: $1, Right: $3}
	}

expr
	: variable '=' expr
	{
		$$ = &ast.AssExpr{Left: []ast.Expr{$1}, Right: []ast.Expr{$3}}
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
	/* TERNARY OPERATOR */
	| expr '?' expr ':' expr
	{
		$$ = &ast.TriOpExpr{Cond: $1, Then: $3, Else: $5}
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
	| common_expr
	{
		$$ = $1
	}

common_expr
	: simp_expr
	{
		$$ = $1
	}
	/* CONCATENATE */
	| common_expr simp_expr %prec CONCAT_OP
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "CAT", Right: $2}
	}

simp_expr
	: non_post_simp_expr
	{
		$$ = $1
	}
	/* ARITHMETIC EXPRESSION */
	| simp_expr '+' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "+", Right: $3}
	}
	| simp_expr '-' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "-", Right: $3}
	}
	| simp_expr '*' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "*", Right: $3}
	}
	| simp_expr '/' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "/", Right: $3}
	}
	| simp_expr '%' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "%", Right: $3}
	}
	/* RELATION EXPRESSION */
	| simp_expr EQEQ simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "==", Right: $3}
	}
	| simp_expr NEQ simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "!=", Right: $3}
	}
	| simp_expr '>' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">", Right: $3}
	}
	| simp_expr GE simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">=", Right: $3}
	}
	| simp_expr '<' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<", Right: $3}
	}
	| simp_expr LE simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<=", Right: $3}
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
	/* COMPOSITE EXPRESSION */
	| simp_expr PLUSPLUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "++", After:true}
	}
	| simp_expr MINUSMINUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "--", After:true}
	}

non_post_simp_expr
	: '!' simp_expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "!", Expr:$2}
	}
	/* FUNCTION CALL */
	| IDENT '(' opt_exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.Literal, SubExprs:$3}
	}
	/*
	| simp_expr '(' opt_exprs ')'
	*/
	| non_post_simp_expr '(' opt_exprs ')'
	{
		$$ = &ast.AnonymousCallExpr{Expr: $1, SubExprs:$3}
	}
	/* FUNCTION DEFINITION */
	| FUNC '(' ident_args ')' '{' opt_stmts '}'
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $6}
	}
	/* ARITHMETIC EXPRESSION */
	| '(' expr ')'
	{
		$$ = &ast.ParentExpr{SubExpr: $2}
	}
	/* COMPOSITE EXPRESSION */
	| PLUSPLUS simp_expr
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "++"}
	}
	| MINUSMINUS simp_expr
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "--"}
	}
	/* LITERAL */
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
	/* var */
	| variable
	{
		$$ = $1
	}
	/* UNARY EXPRESSION */
	| '+' simp_expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "+", Expr:$2}
	}
	| '-' simp_expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "-", Expr:$2}
	}
	/*
	*/

variables
	: variable
	{
		$$ = []ast.Expr{$1}
	}
	| variables ',' opt_term variable
	{
		$$ = append($1,$4)
	}

variable
	/*
	: simp_expr '[' exprs ']'
	: variable '[' exprs ']'
	: non_post_simp_expr '[' exprs ']'
	*/
	: non_post_simp_expr '[' exprs ']'
	{
		$$ = &ast.ItemExpr{Expr: $1, Index:$3}
	}
	/*
	*/
	| IDENT
	{
		$$ = &ast.IdentExpr{Literal: $1.Literal}
	}
	| '$' non_post_simp_expr
	{
		$$ = &ast.FieldExpr{Expr: $2}
	}
	/*
	*/

ident_args
	: /* empty */
	{
		$$ = []string{}
	}
	| IDENT
	{
		$$ = []string{$1.Literal}
	}
	| ident_args ',' opt_term IDENT
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
	: semi
	| term semi

semi
	: ';'  /* go/scanner return semi when EOL */

/*
opt_nls
	: 
	| nls

nls
	: '\n'
	| nls '\n'
*/

%%
