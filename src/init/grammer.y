%{

package init

import (
	//"fmt"
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
	"Compiler/src/ast/statement/function"
	"Compiler/src/symbolTable"
	"Compiler/src/symbolTable/symbol"
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/rvalue"
	stmt "Compiler/src/ast/statement"
	"Compiler/src/ast/statement/while"
	"Compiler/src/symbolTable"
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

%token <string_value> STRING INT DOUBLE CHAR BOOL VOID

%token IF ELSE WHILE FOR PRINT EOL

%token <string_value> EQ NE LT LE GT GE ASSIGN COMMA
%token <string_value> ADD SUB MUL DIV NOT
%token <string_value> LBRACKET RBRACKET LPARENTHESIS RPARENTHESIS LBRACE RBRACE

%left EQ NE LT LE GT GE
%left ADD SUB
%left MUL DIV

%type <node> expression statement program assign block stmtList funcDefine args
%type <node> atomExpression unaryExpression binaryOrAtomExpression refExpression
%type <node> assignStatement printStatement declareStatement whileStatement
//%type <string_value> unaryOperator


%%

program
	: {
		result = compound.CreateCompoundStmt()
		$$ = result
	}
    	| program statement EOL{
    		v1 := ($1).(*compound.CompoundStmt)
    		v2 := ($2).(stmt.Stmt)
		compound.AddStmt(v1, v2)
    	}
	;

assign
	: ASSIGN expression {
    		$$=$2
    	}
    	;

declareStatement
	: TYPE IDENTIFY {
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

assignStatement
	: refExpression assign {
    		v1 := ($1).(lvalue.LValue)
    		v2 := ($2).(rvalue.RValue)
		$$ = assign.CreateAssignStmt(v1, v2)
    	}
    	;

printStatement
	: PRINT LPARENTHESIS expression RPARENTHESIS {
		v3 := ($3).(rvalue.RValue)
		$$ = print.CreatePrintStmt(v3)
    	}
	;

whileStatement
	: WHILE expression block {
		symbolTable.PushFrame()
		v2 := ($2).(rvalue.RValue)
		v3 := ($3).(*compound.CompoundStmt)
		$$ = while.CreateWhileStmt(v2, v3)
		symbolTable.PopFrame()
	}

//returnStatement TODO

statement
	: assignStatement
	| printStatement
	| declareStatement
	| block
	| whileStatement
	| funcDefine
	;


atomExpression
	: INT_LITERAL {
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

refExpression
	: IDENTIFY {
		sb := symbolTable.GetSymbol($1)
		if sb == nil{
			panic("[*] Error: Undefined var")
		}
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
    	| binaryOrAtomExpression EQ unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(EQ, v1, v3)
	}
    	;

expression
	: binaryOrAtomExpression {
		$$=$1
	}
	;

block
	: LBRACE RBRACE {
		//fmt.Println("None stmts func")
		$$ = nil
	}
	| LBRACE stmtList RBRACE {
		$$=$2
	}

stmtList
	: stmtList statement {
		v1 := ($1).(*compound.CompoundStmt)
		v2 := ($2).(stmt.Stmt)
		compound.AddStmt(v1, v2)
	}
	| {
		$$ = compound.CreateCompoundStmt()
	}

//funcCallStmt
//	:

// args处返回CreateDeclareStmt切片，block处返回函数的CompoundStmt
funcDefine
	: TYPE IDENTIFY LPARENTHESIS {symbolTable.PushFrame()} args RPARENTHESIS block {
		//CreateFunc()
		v1 := $1				 // function return type
		/*
			function name, add to symbolTable
                	(作用域？全局符号表第几级？)
                	update：目前仍全局符号栈里了
                */
		v2 := ($2)
		if ($7) == nil {
			$$ = function.CreateFuncDefine(v1, v2, nil, nil)
		} else {
			v7 := ($7).(*compound.CompoundStmt)	 // function stmts
			if ($5) != nil{
				v5 := ($5).([]*declare.DeclareStmt) // function args
				$$ = function.CreateFuncDefine(v1, v2, v5, v7)
			} else {
				$$ = function.CreateFuncDefine(v1, v2, nil, v7)
			}
		}
		symbolTable.PopFrame()
	}
	;

args
	: {
		//fmt.Println("None args func")
		$$ = nil
	}
	| declareStatement {
		//fmt.Println("args branch")
		var declareStmts []*declare.DeclareStmt
		v1 := ($1).(declare.DeclareStmt)
		$$ = append(declareStmts, &v1)
	}
	| args COMMA declareStatement {
		v1 := ($1).([]*declare.DeclareStmt)
		v3 := ($3).(declare.DeclareStmt)
		$$ = append(v1, &v3)
		//fmt.Println("Another args")
	}

%%