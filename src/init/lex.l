
%option noyywrap

%{
#include "token.h"
%}

%%

"void"                      {return VOID;}
"char"                      {return CHAR;}
"bool"                      {return BOOL;}
"int"                       {return INT;}
"double"                    {return DOUBLE;}
"string"                    {return STRING;}
"if"                        {return IF;}
"else"                      {return ELSE;}
"while"                     {return WHILE;}
"print"                     {return PRINT;}
([-])?[0-9]+                {return INT_LITERAL;}
([-])?[0-9]+\.[0-9]*        {return DOUBLE_LITERAL;}
[a-zA-Z][0-9a-zA-Z_]*       {return IDENTIFY;}
"=="                        {return EQ;}
"!="                        {return NE;}
"<"                         {return LT;}
"<="                        {return LE;}
">"                         {return GT;}
">="                        {return GE;}
"+"                         {return ADD;}
"-"                         {return SUB;}
"*"                         {return MUL;}
"/"                         {return DIV;}
"["                         {return LBRACKET;}
"]"                         {return RBRACKET;}
"("                         {return LPARENTHESIS;}
")"                         {return RPARENTHESIS;}
"{"                         {return LBRACE;}
"}"                         {return RBRACE;}
"\n"                        {return EOL;}
"="                         {return ASSIGN;}
[ \t]                       {}
"!"                         {return NOT;}
","                         {return COMMA;}
%%