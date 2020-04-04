package symbolTable

import (
	"Compiler/src/symbolTable/symbol"
)

var GlobalSymbolTableStack = SymbolTableStack{0, nil}
var NextFrameId = 0

type SymbolTableFrame struct {
	Length 		int
	Symbols 	[]*symbol.Symbol
	NamespaceId	int
}

type SymbolTableStack struct {
	Length 		int
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
	for _, f := range GlobalSymbolTableStack.Frames {
		for _, s := range f.Symbols {
			if s.Name == name {
				return s
			}
		}
	}
	return nil
}