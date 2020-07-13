package binOperateResult

import (
	"Compiler/src/ast/expression/lvalue/reference"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/ast/expression/rvalue/literal"
	"Compiler/src/global"
	"Compiler/src/symbolTable"
	"fmt"
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
			fmt.Print("fdiv")
		} else {
			fmt.Print("sdiv")
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
		panic("unknown op")
	}
	fmt.Printf(" %s ", symbolTable.TypeString(b.Type))
	//if symbolTable.TypeString(b.Type) == "double" {
	//	ir, _ := strconv.ParseFloat(b.RHS.RvalueIR(),64)
	//	IRS := strconv.FormatFloat(ir,'f',-1,64)
	//	fmt.Printf("%s, %s\n", b.LHS.RvalueIR(), IRS)
	//} else {
	//	fmt.Printf("%s, %s\n", b.LHS.RvalueIR(), b.RHS.RvalueIR())
	//}
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
	var NRHS literal.DoubleLiteral
	if LHS.GetType() != RHS.GetType() {
		x := LHS.(reference.VariableReference)
		if x.Type == global.DOUBLE_LITERAL {
			switch RHS.(type) {
			case literal.IntLiteral:
				tmp := RHS.(literal.IntLiteral)
				NRHS.Type = global.DOUBLE_LITERAL
				NRHS.Value = float64(tmp.Value)
				result.RHS = NRHS
			default:
				panic("operator type not match")
			}
		}
	} else {
		result.RHS = RHS // 避免double a = b - 1;
	}
	result.Type = LHS.GetType()
	result.OperatorId = OperateId
	result.LHS = LHS

	rvalue.NextTmpReg++
	result.TmpRegId = rvalue.NextTmpReg
	return result
}

