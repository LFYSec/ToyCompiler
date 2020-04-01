%{

package main

import (
	"fmt"
	"strconv"
)

type SymbolType int

%}

%union {
	int_value 	int
	double_value 	float64
	string_value 	string
   	node 		*ASTNode;
   	symbolType 	SymbolType
}

%token <symbolType> 	TYPE
%token <string_value> 	IDENTIFY
%token <int_value> 	INT_LITERAL
%token <double_value> 	DOUBLE_LITERAL

%token <string_value> STRING INT DOUBLE CHAR BOOL

%token IF ELSE WHILE FOR PRINT EOL

%token <string_value> EQ NE LT LE GT GE ASSIGN
%token <string_value> ADD SUB MUL DIV NOT
%token <string_value> LBRACKET RBRACKET

%left EQ NE LT LE GT GE
%left ADD SUB
%left MUL DIV

%type <string_value> expression statement
%type <string_value> atomExpression unaryExpression binaryOrAtomExpression
%type <string_value> assignStatement printStatement declareStatement
%type <string_value> unaryOperator assign


%%

program:
    program statement EOL          		{fmt.Println("statement: ", $2)}
    |				        	{fmt.Println("start")}
    ;

assign:
    ASSIGN expression              {
    	$$=$2
    	fmt.Println("assign: ",$2)
    }
    ;

declareStatement:
    TYPE IDENTIFY                     {
    		$$ = strconv.Itoa(int($1))
           	fmt.Println("TYPE: ",$1)
        }
    | TYPE IDENTIFY assign           {
		fmt.Println("TYPE: ",$1)
        }
    ;

assignStatement:
    	IDENTIFY assign {
            $$=$2
    	}
    	;

printStatement:
    	PRINT '(' expression ')'           {fmt.Println("PRINT: ",$3)}
	;

statement:
	assignStatement
	| printStatement
	| declareStatement
	;


atomExpression:
    INT_LITERAL                             {
	$$ = strconv.Itoa($1)
    	fmt.Println("INT_LITERAL: ", $1)
    }
    | DOUBLE_LITERAL                        {
    	$$=strconv.FormatFloat($1, 'E', -1, 64)
    	fmt.Println("DOU_LITERAL: ",$1)
    }
    | IDENTIFY                		    {
    	$$=$1
	fmt.Println("IDENTIFY: ",$1)
    }
    | '(' expression ')'                   {$$=$2}
    ;

unaryOperator:
    ADD
    | SUB
    | NOT
    ;

unaryExpression:
    atomExpression                          {$$=$1}
    | unaryOperator atomExpression          {$$=$2}
;

binaryOrAtomExpression:
    unaryExpression                                     {$$=$1}
    | binaryOrAtomExpression '+' unaryExpression        {$$=$1+$3}
    ;

expression:
    binaryOrAtomExpression                           {
	$$=$1
    }
    ;


%%