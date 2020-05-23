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

func (b BinOperateResult) GeneRVCode() {
	b.LHS.GeneRVCode()
	b.RHS.GeneRVCode()
	irS := b.RvalueIR()
	fmt.Print(irS, " = ")
	switch b.OperatorId {
	case global.ADD:
		if b.Type == global.DOUBLE_LITERAL {
			fmt.Print("fadd")
		} else {
			fmt.Print("add nsw")
		}
		break
	case global.SUB:
		if b.Type == global.DOUBLE_LITERAL {
			fmt.Print("fsub")
		} else {
			fmt.Print("sub nsw")
		}
		break
	case global.MUL:
		if b.Type == global.DOUBLE_LITERAL {
			fmt.Print("fmul")
		} else {
			fmt.Print("mul nsw")
		}
		break
	case global.DIV:
		if b.Type == global.DOUBLE_LITERAL {
			fmt.Print("fdif")
		} else {
			fmt.Print("dif nsw")
		}
		break
	case global.GT:
		switch b.LHS.GetType() {
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
		switch b.LHS.GetType() {
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
		switch b.LHS.GetType() {
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
		switch b.LHS.GetType() {
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
		switch b.LHS.GetType() {
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
		switch b.LHS.GetType() {
		case global.INT_LITERAL:
			fmt.Print("icmp ne")
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
	fmt.Printf(" %s ", symbolTable.TypeString(b.Type))
	fmt.Printf("%s, %s\n", b.LHS.RvalueIR(), b.RHS.RvalueIR())
}

func (b BinOperateResult) RvalueIR() string {
	result := fmt.Sprintf("%%temp_%d", b.TmpRegId)
	return result
}

func (b BinOperateResult) GetType() global.SymbolType {
	return b.Type
}


func CreateBinOperateResult(OperateId int, LHS,RHS rvalue.RValue) BinOperateResult {
	var result BinOperateResult
	if LHS.GetType() != RHS.GetType() {
		os.Exit(1)
	}
	//switch OperateId {
	//case global.GT:
	//case global.LT:
	//	result.Type = global.BOOL
	//	break
	//default:
	//	result.Type = LHS.GetType()
	//	break
	//}
	result.Type = LHS.GetType()
	result.OperatorId = OperateId
	result.LHS = LHS
	result.RHS = RHS
	rvalue.NextTmpReg++
	result.TmpRegId = rvalue.NextTmpReg
	return result
}

