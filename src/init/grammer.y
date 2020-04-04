%{

package init

import (
	"Compiler/src/ast/statement/compound"
	"Compiler/src/ast"
	"Compiler/src/ast/expression/rvalue/binOperateResult"
	"Compiler/src/ast/statement/compound"
	"Compiler/src/ast/statement/declare"
	"Compiler/src/ast/statement/print"
	"Compiler/src/symbolTable"
	"Compiler/src/symbolTable/symbol"
	"Compiler/src/global"
	"Compiler/src/ast"
	"Compiler/src/ast/expression/lvalue/variableReference"
	"Compiler/src/ast/expression/rvalue/doubleLiteral"
	"Compiler/src/ast/expression/rvalue/intLiteral"
	"Compiler/src/ast/statement/assign"
	"Compiler/src/symbolTable"
	"Compiler/src/symbolTable/symbol"
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/rvalue"
	stmt "Compiler/src/ast/statement"
)

var result *compound.CompoundStmt

%}

%union {
	int_value 	int
	double_value 	float64
	string_value 	string
   	node 		ast.Node
   	symbolType 	global.SymbolType
}

%token <symbolType> 	TYPE
%token <string_value> 	IDENTIFY
%token <int_value> 	INT_LITERAL
%token <double_value> 	DOUBLE_LITERAL
%token ARRAY

%token <string_value> STRING INT DOUBLE CHAR BOOL

%token IF ELSE WHILE FOR PRINT EOL

%token <string_value> EQ NE LT LE GT GE ASSIGN
%token <string_value> ADD SUB MUL DIV NOT
%token <string_value> LBRACKET RBRACKET LPARENTHESIS RPARENTHESIS

%left EQ NE LT LE GT GE
%left ADD SUB
%left MUL DIV

%type <node> expression statement program assign
%type <node> atomExpression unaryExpression binaryOrAtomExpression refExpression
%type <node> assignStatement printStatement declareStatement
//%type <string_value> unaryOperator


%%

program
	: {
		result = compound.CreateCompoundStmt()
		$$ = result
	}
    	| program statement EOL	{
    		v1 := ($1).(*compound.CompoundStmt)
    		v2 := ($2).(stmt.Stmt)
		compound.AddStmt(v1, v2)
    	}
	;

assign:
    	ASSIGN expression {
    		$$=$2
    	}
    	;

declareStatement:
    	TYPE IDENTIFY {
		s := symbol.CreateSymbol(true, $1, $2)
		symbolTable.AddSymbol(s)
		$$ = declare.CreateDeclareStmt(s, assign.AssignStmt{})
        }
    	| TYPE IDENTIFY assign {
		s := symbol.CreateSymbol(true, $1, $2)
		symbolTable.AddSymbol(s)
		ref := variableReference.CreateVariableReference(s)
		v := ($3).(rvalue.RValue)
		a := assign.CreateAssignStmt(ref, v)
		$$ = declare.CreateDeclareStmt(s, a)
        }
    	;

assignStatement:
    	refExpression assign {
    		v1 := ($1).(lvalue.LValue)
    		v2 := ($2).(rvalue.RValue)
		$$ = assign.CreateAssignStmt(v1, v2)
    	}
    	;

printStatement:
    	PRINT LPARENTHESIS expression RPARENTHESIS {
    		v3 := ($3).(rvalue.RValue)
    		$$ = print.CreatePrintStmt(v3)
    	}
	;

statement:
	assignStatement
	| printStatement
	| declareStatement
	;


atomExpression:
    	INT_LITERAL {
		$$ = intLiteral.CreateIntLiteral($1)
    	}
    	| DOUBLE_LITERAL {
		$$ = doubleLiteral.CreateDoubleLiteral($1)
    	}
    	| refExpression {
    		$$=$1
    	}
    	| LPARENTHESIS expression RPARENTHESIS {
    		$$=$2
    	}
    	;

refExpression:
	IDENTIFY {
		sb := symbolTable.GetSymbol($1)
                $$ = variableReference.CreateVariableReference(sb)
	}

unaryOperator
	: ADD
	| SUB
	| NOT
	;

unaryExpression
	: atomExpression 				{$$=$1}
	| unaryOperator atomExpression          	{$$=$2}
	;

binaryOrAtomExpression
	: unaryExpression                       	{$$=$1}
    	| binaryOrAtomExpression ADD unaryExpression {
    		v1 := ($1).(rvalue.RValue)
    		v3 := ($3).(rvalue.RValue)
    		$$ = binOperateResult.CreateBinOperateResult(ADD, v1, v3)
    	}
    	;

expression
	: binaryOrAtomExpression {
		$$=$1
	}
	;


%%