%{

package init

import (
	"fmt"
	"Compiler/src/ast"
	"Compiler/src/global"
	"Compiler/src/ast/expression/lvalue/reference"
	"Compiler/src/ast/expression/rvalue/literal"
	"Compiler/src/ast/expression/rvalue/binOperateResult"
	"Compiler/src/ast/statement/function"
	"Compiler/src/symbolTable"
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/rvalue"
	stmt "Compiler/src/ast/statement"
)

var result *stmt.CompoundStmt

%}

%union {
	int_value 	int
	double_value 	float64
	string_value 	string
	char_value	uint8
   	node 		ast.Node
   	symbolType 	global.SymbolType
}

%token <symbolType> 	TYPE
%token <string_value> 	IDENTIFY
%token <int_value> 	INT_LITERAL
%token <double_value> 	DOUBLE_LITERAL
%token <char_value> 	CHAR_LITERAL
%token ARRAY

%token <string_value> STRING INT DOUBLE CHAR BOOL VOID

%token IF ELSE WHILE FOR PRINT EOL RETURN EOS SCAN

%token <string_value> EQ NE LT LE GT GE ASSIGN COMMA
%token <string_value> ADD SUB MUL DIV NOT
%token <string_value> LBRACKET RBRACKET LPARENTHESIS RPARENTHESIS LBRACE RBRACE

%left EQ NE LT LE GT GE
%left ADD SUB
%left MUL DIV

%type <node> expression statement program assign block stmtList defargs callargs
%type <node> atomExpression binaryOrAtomExpression refExpression unaryExpression
%type <node> assignStatement printStatement declareStatement whileStatement scanStatement
%type <node> returnStatement funcDefineStmt funCallStmt assignFunCallStmt ifStatement


%%

program
	: {
		result = stmt.CreateCompoundStmt()
		$$ = result
	}
    	| program statement {
    		v1 := ($1).(*stmt.CompoundStmt)
    		v2 := ($2).(stmt.Stmt)
		stmt.AddStmt(v1, v2)
    	}
	;

declareStatement
	: TYPE IDENTIFY EOS {
		s := symbolTable.CreateSymbol(true, $1, $2, 0)
		symbolTable.AddSymbol(s)
		$$ = stmt.CreateDeclareStmt(s, stmt.AssignStmt{})
        }
    	| TYPE IDENTIFY assign EOS {
		s := symbolTable.CreateSymbol(true, $1, $2, 0)
		symbolTable.AddSymbol(s)
		ref := reference.CreateVariableReference(s)
		v := ($3).(rvalue.RValue)
		a := stmt.CreateAssignStmt(ref, v)
		$$ = stmt.CreateDeclareStmt(s, a)
        }
        // array每个元素当成一个变量存，然后整个array通过同一个array_name的符号表来找
        | TYPE IDENTIFY LBRACKET INT_LITERAL RBRACKET EOS {
		s := symbolTable.CreateSymbol(true, $1, $2, $4)
		symbolTable.AddSymbol(s)
		$$ = stmt.CreateDeclareStmt(s, stmt.AssignStmt{})
        }
    	;

assign
	: ASSIGN expression {
    		$$=$2
    	}
    	;

assignStatement
	: refExpression assign EOS {
    		v1 := ($1).(lvalue.LValue)
    		v2 := ($2).(rvalue.RValue)
		$$ = stmt.CreateAssignStmt(v1, v2)
    	}
    	;

printStatement
	: PRINT LPARENTHESIS expression RPARENTHESIS EOS {
		v3 := ($3).(rvalue.RValue)
		$$ = stmt.CreatePrintStmt(v3)
    	}
	;

scanStatement
	: SCAN LPARENTHESIS expression RPARENTHESIS EOS {
		v3 := ($3).(rvalue.RValue)
		$$ = stmt.CreateScanStmt(v3)
    	}
	;

ifStatement
	: IF expression block {
		symbolTable.PushFrame()
		v2 := ($2).(rvalue.RValue)
		v3 := ($3).(*stmt.CompoundStmt)
		$$ = stmt.CreateIfStmt(v2, v3, nil)
		symbolTable.PopFrame()
	}
	| IF expression block ELSE block {
		symbolTable.PushFrame()
		v2 := ($2).(rvalue.RValue)
		v3 := ($3).(*stmt.CompoundStmt)
		v5 := ($5).(*stmt.CompoundStmt)
		$$ = stmt.CreateIfStmt(v2, v3, v5)
		symbolTable.PopFrame()
	}
	;

whileStatement
	: WHILE expression block {
		symbolTable.PushFrame()
		v2 := ($2).(rvalue.RValue)
		v3 := ($3).(*stmt.CompoundStmt)
		$$ = stmt.CreateWhileStmt(v2, v3)
		symbolTable.PopFrame()
	}
	;

returnStatement
	: RETURN expression EOS {
		v2 := ($2).(rvalue.RValue)
		$$ = stmt.CreateRetStmt(v2)
	}
	;

statement
	: returnStatement
	| assignStatement
	| printStatement
	| declareStatement
	| block
	| whileStatement
	| funcDefineStmt
	| funCallStmt
	| assignFunCallStmt
	| ifStatement
	| scanStatement
	;

unaryExpression
	: atomExpression {
		$$=$1;
	}
	;

atomExpression
	: INT_LITERAL {
		$$ = literal.CreateIntLiteral($1)
    	}
    	| DOUBLE_LITERAL {
		$$ = literal.CreateDoubleLiteral($1)
    	}
    	| refExpression {
    		$$=$1
    	}
    	| LPARENTHESIS expression RPARENTHESIS {
    		$$=$2
    	}
    	| CHAR_LITERAL {
    		$$ = literal.CreateCharLiteral($1)
    	}
    	;

refExpression
	: IDENTIFY {
		sb := symbolTable.GetVarSymbol($1)
		if sb == nil {
			panic("[*] Error: Undefined variable: "+string($1))
		}
                $$ = reference.CreateVariableReference(sb)
	}
	| IDENTIFY LBRACKET expression RBRACKET {
		sb := symbolTable.GetVarSymbol($1)
		if sb == nil {
			panic("[*] Error: Undefined array: "+string($1))
		}
		v3 := ($3).(rvalue.RValue)
		$$ = reference.CreateArrayReference(sb, v3)
	}
	;

binaryOrAtomExpression
	: unaryExpression { // 最低优先级，防止比如print(a-1)识别为a,-1
//		fmt.Println("test0")
		$$=$1
	}
	| binaryOrAtomExpression ADD unaryExpression {
    		v1 := ($1).(rvalue.RValue)
    		v3 := ($3).(rvalue.RValue)
    		$$ = binOperateResult.CreateBinOperateResult(ADD, v1, v3)
    	}
    	| binaryOrAtomExpression SUB unaryExpression {
//    		fmt.Println("test1")
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(SUB, v1, v3)
	}
	| binaryOrAtomExpression MUL unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(MUL, v1, v3)
	}
	| binaryOrAtomExpression DIV unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(DIV, v1, v3)
	}
	| binaryOrAtomExpression GT unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(GT, v1, v3)
	}
	| binaryOrAtomExpression LT unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(LT, v1, v3)
	}
	| binaryOrAtomExpression LE unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(LE, v1, v3)
	}
	| binaryOrAtomExpression GE unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(GE, v1, v3)
	}
	| binaryOrAtomExpression NE unaryExpression {
		v1 := ($1).(rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = binOperateResult.CreateBinOperateResult(NE, v1, v3)
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
	: LBRACE stmtList RBRACE {
		$$=$2
	}
	;

stmtList
	: stmtList statement {
		v1 := ($1).(*stmt.CompoundStmt)
		v2 := ($2).(stmt.Stmt)
		stmt.AddStmt(v1, v2)
	}
	| {
		$$ = stmt.CreateCompoundStmt()
	}
	;

// funCallStmt返回的不能是一个Stmt节点，因为要将函数返回值赋给assign语句，而assign语句需要RValue。
// 解决方案是funCallStmt单独加到CompoundStmt中，并且还要返回callStmt的计算结果
funCallStmt
	: IDENTIFY LPARENTHESIS callargs RPARENTHESIS {
		v1 := $1
		if ($3) != nil {
			v3 := ($3).([]rvalue.RValue)
			$$ = function.CreateFuncCallStmt(v1, v3)
		} else {
			$$ = function.CreateFuncCallStmt(v1, nil)
		}
	}
	| funCallStmt EOS {
		$$ = $1
	}
	;

assignFunCallStmt
    	: refExpression ASSIGN funCallStmt EOS {
		v1 := ($1).(lvalue.LValue)
		v3 := ($3).(stmt.Stmt)
		$$ = function.CreateAssignFunCallStmt(v1, v3)
    	}
    	// TODO Fix int a = factorial(3);
    	| TYPE IDENTIFY ASSIGN funCallStmt EOS {
		s := symbolTable.CreateSymbol(true, $1, $2, 0)
		symbolTable.AddSymbol(s)
		ref := reference.CreateVariableReference(s)
		v4 := ($4).(stmt.Stmt)
		$$ = function.CreateAssignFunCallStmt(ref, v4)
    	}
    	;

// args处返回CreateDeclareStmt切片，block处返回函数的CompoundStmt
funcDefineStmt
	: TYPE IDENTIFY LPARENTHESIS {symbolTable.PushFrame()} defargs RPARENTHESIS {
		if ($5) == nil{
			function.AddFunc($1,$2,nil)
		} else {
			v5 := ($5).([]*stmt.DeclareStmt)
			function.AddFunc($1,$2,v5)
		}
	} block {
		/*
			function name, add to symbolTable
                	(作用域？全局符号表第几级？)
                	update：目前仍全局符号栈里了
                */
		v2 := $2
		v8 := ($8).(*stmt.CompoundStmt)
		$$ = function.CreateFuncDefine(v2, v8)
		symbolTable.PopFrame()
	}
	;

defargs
	: {
		$$ = nil
	}
	| defargs TYPE IDENTIFY {
		s := symbolTable.CreateSymbol(true, $2, $3, 0)
		symbolTable.AddSymbol(s)
		v1 := stmt.CreateDeclareStmt(s, stmt.AssignStmt{})
		var declareStmts []*stmt.DeclareStmt
		$$ = append(declareStmts, &v1)
	}
	| defargs COMMA TYPE IDENTIFY {
		s := symbolTable.CreateSymbol(true, $3, $4, 0)
		symbolTable.AddSymbol(s)
		v3 := stmt.CreateDeclareStmt(s, stmt.AssignStmt{})

		v1 := ($1).([]*stmt.DeclareStmt)
		$$ = append(v1, &v3)
	}
	;

callargs
	: {
		$$ = nil
	}
	| callargs expression {
		var rv []rvalue.RValue
		v2 := ($2).(rvalue.RValue)
		$$ = append(rv, v2)
	}
	| callargs COMMA expression {
		v1 := ($1).([]rvalue.RValue)
		v3 := ($3).(rvalue.RValue)
		$$ = append(v1, v3)
		fmt.Println("Another callargs")
	}
	;

%%