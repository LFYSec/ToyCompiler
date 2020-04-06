package binOperateResult

import (
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"Compiler/src/symbolTable"
	"fmt"
	"os"
)


type BinOperateResult struct {
	rvalue.RValue
	TmpRegId  	int
	OperatorId 	int
	Type 		global.SymbolType
	LHS        	rvalue.RValue
	RHS        	rvalue.RValue
}

func (binOperateResult BinOperateResult) GeneRVCode() {
	binOperateResult.LHS.GeneRVCode()
	binOperateResult.RHS.GeneRVCode()
	irS := binOperateResult.RvalueIR()
	fmt.Print(irS, " = ")
	switch binOperateResult.OperatorId {
	case global.ADD:
		if binOperateResult.Type == global.DOUBLE_LITERAL {
			fmt.Print("fadd")
		} else {
			fmt.Print("add nsw")
		}
		break
	case global.SUB:
		if binOperateResult.Type == global.DOUBLE_LITERAL {
			fmt.Print("fsub")
		} else {
			fmt.Print("sub nsw")
		}
		break
	case global.MUL:
		if binOperateResult.Type == global.DOUBLE_LITERAL {
			fmt.Print("fmul")
		} else {
			fmt.Print("mul nsw")
		}
		break
	case global.DIV:
		if binOperateResult.Type == global.DOUBLE_LITERAL {
			fmt.Print("fdif")
		} else {
			fmt.Print("dif nsw")
		}
		break
	case global.GT:
		switch binOperateResult.LHS.GetType() {
		case global.INT_LITERAL:
			fmt.Print("icmp sgt")
			break
		case global.DOUBLE_LITERAL:
			fmt.Print("fcmp ogt")
			break
		default:
			break
		}
	case global.LT:
		switch binOperateResult.LHS.GetType() {
		case global.INT_LITERAL:
			fmt.Print("icmp slt")
			break
		case global.DOUBLE_LITERAL:
			fmt.Print("fcmp olt")
			break
		default:
			break
		}
	case global.LE:
		switch binOperateResult.LHS.GetType() {
		case global.INT_LITERAL:
			fmt.Print("icmp sle")
			break
		case global.DOUBLE_LITERAL:
			fmt.Print("fcmp ole")
			break
		default:
			break
		}
	case global.GE:
		switch binOperateResult.LHS.GetType() {
		case global.INT_LITERAL:
			fmt.Print("icmp sge")
			break
		case global.DOUBLE_LITERAL:
			fmt.Print("fcmp oge")
			break
		default:
			break
		}
	case global.EQ:
		switch binOperateResult.LHS.GetType() {
		case global.INT_LITERAL:
			fmt.Print("icmp eq")
			break
		case global.DOUBLE_LITERAL:
			fmt.Print("fcmp oeq")
			break
		default:
			break
		}
	case global.NE:
		switch binOperateResult.LHS.GetType() {
		case global.INT_LITERAL:
			fmt.Print("icmp.ne")
			break
		case global.DOUBLE_LITERAL:
			fmt.Print("fcmp one")
			break
		default:
			break
		}
	default:
		os.Exit(1)
	}
	fmt.Printf(" %s ", symbolTable.TypeString(binOperateResult.Type))
	fmt.Printf("%s, %s\n", binOperateResult.LHS.RvalueIR(), binOperateResult.RHS.RvalueIR())
}

func (binOperateResult BinOperateResult) RvalueIR() string {
	result := fmt.Sprintf("%%temp_%d", binOperateResult.TmpRegId)
	return result
}

func (binOperateResult BinOperateResult) GetType() global.SymbolType {
	return binOperateResult.Type
}


func CreateBinOperateResult(OperateId int, LHS,RHS rvalue.RValue) BinOperateResult {
	var result BinOperateResult
	if LHS.GetType() != RHS.GetType() {
		os.Exit(1)
	}
	switch OperateId {
	case global.GT:
	case global.LT:
		result.Type = global.BOOL
		break
	default:
		result.Type = LHS.GetType()
		break
	}
	result.OperatorId = OperateId
	result.LHS = LHS
	result.RHS = RHS
	rvalue.NextTmpReg++
	result.TmpRegId = rvalue.NextTmpReg
	return result
}

