package symbolTable

import (
	"Compiler/src/symbolTable/symbol"
)

var GlobalSymbolTableStack = SymbolTableStack{0, nil, nil}
var NextFrameId = 0

type SymbolTableFrame struct {
	Length 		int
	Symbols 	[]*symbol.Symbol
	NamespaceId	int
}

type SymbolTableStack struct {
	Length 		int
	Funcs		[]*symbol.FuncSymbol
	Frames	 	[]*SymbolTableFrame
}

func PushFrame() {
	NextFrameId++
	newFrame := SymbolTableFrame{
		Length:      0,
		Symbols:     nil,
		NamespaceId: NextFrameId,
	}
	GlobalSymbolTableStack.Frames = append(GlobalSymbolTableStack.Frames, &newFrame)
	GlobalSymbolTableStack.Length++
}

func PopFrame() {
	GlobalSymbolTableStack.Length--
	GlobalSymbolTableStack.Frames = GlobalSymbolTableStack.Frames[:len(GlobalSymbolTableStack.Frames)-1]
}

func AddSymbol(s *symbol.Symbol) {
	frame := GlobalSymbolTableStack.Frames[GlobalSymbolTableStack.Length-1]
	frame.Symbols = append(frame.Symbols, s)
	frame.Length++
}

func GetSymbol(name string) *symbol.Symbol {
	for i := GlobalSymbolTableStack.Length-1; i>=0 ;i-- {
		f := GlobalSymbolTableStack.Frames[i] // 通过每一层block先push_frame实现作用域。栈结构，后push的优先级高
		for _, s := range f.Symbols {
			if s.Name == name {
				return s
			}
		}
	}
	return nil
}

func AddFunc(f *symbol.FuncSymbol) {
	_ = append(GlobalSymbolTableStack.Funcs, f)
}