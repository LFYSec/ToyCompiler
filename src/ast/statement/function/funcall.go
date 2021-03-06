package function

import (
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/lvalue/reference"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/ast/expression/rvalue/binOperateResult"
	"Compiler/src/ast/expression/rvalue/literal"
	stmt "Compiler/src/ast/statement"
	st "Compiler/src/symbolTable"
	"fmt"
)

type Funcall struct {
	stmt.Stmt
	FuncName 			string
	FuncArgsValue		[]rvalue.RValue
	LV 					lvalue.LValue
}

type AssignFuncall struct {
	Func	 			stmt.Stmt
	LVariable			lvalue.LValue
}

func (f Funcall) GeneCode() {
	funcArgs := ""
	argFormat := "%s %s"
	//TODO fix type error
	count := 0
	for _, i := range f.FuncArgsValue {
		if count != 0 {
			argFormat += ","
		}
		count ++
		switch i.(type) {
		case literal.IntLiteral:
			arg := i.(literal.IntLiteral)
			funcArgs += fmt.Sprintf(argFormat, st.TypeString(arg.Type), arg.RvalueIR())
		case binOperateResult.BinOperateResult:
			arg := i.(binOperateResult.BinOperateResult)
			arg.GeneRVCode()
			funcArgs += fmt.Sprintf(argFormat, st.TypeString(arg.Type), arg.RvalueIR())

		}
	}
	fmt.Printf("%s = call %s @%s(%s)\n",
		f.LV.LvalueIR(),
		st.TypeString(st.GetFuncSymbol(f.FuncName).ReturnType),
		f.FuncName,
		funcArgs,
	)
}

func (a AssignFuncall) GeneCode() {
	funCall := a.Func.(Funcall)
	funCall.GeneCode()
	RVIR := funCall.LV.LvalueIR()
	LVIR := a.LVariable.LvalueIR()
	fmt.Printf("store %s %s, %s* %s\n",
		st.TypeString(funCall.LV.GetType()),
		RVIR,
		st.TypeString(a.LVariable.GetType()),
		LVIR,
	)
}

// 返回Funcall和RValue，Funcall加入compoundStmt，LValue用于返回值后的assign语句，如果没有assign，比如空返回值，则nil
func CreateFuncCallStmt(name string, rvs []rvalue.RValue) Funcall {
	var funcall Funcall
	funcall.FuncName = name
	for _, i := range rvs {
		funcall.FuncArgsValue = append(funcall.FuncArgsValue, i)
	}
	funcType := st.GetFuncSymbol(name)
	s := st.CreateSymbol(true, funcType.ReturnType, "fvar_"+name, 0)
	st.AddSymbol(s)
	ref := reference.CreateVariableReference(s)
	funcall.LV = ref
	return funcall
}

func CreateAssignFunCallStmt(lhs lvalue.LValue, stmt stmt.Stmt) AssignFuncall {
	var assignFunc AssignFuncall
	assignFunc.Func = stmt
	LV := (lhs).(reference.VariableReference)
	sb := st.GetVarSymbol(LV.Variable.Name)
	if sb == nil {
		panic("[*] Error: Undefined variable: " + LV.Variable.Name)
	}
	assignFunc.LVariable = lhs
	return assignFunc
}