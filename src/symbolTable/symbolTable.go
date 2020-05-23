package symbolTable

var GlobalSymbolTableStack = Stack{0, nil, nil}
var NextFrameId = 0

type Frame struct {
	Length 		int
	Symbols 	[]*Symbol
	NamespaceId	int
}

type Stack struct {
	Length 		int
	Funcs		[]*FuncSymbol
	Frames	 	[]*Frame
}

func PushFrame() {
	NextFrameId++
	newFrame := Frame{
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

func AddSymbol(s *Symbol) {
	frame := GlobalSymbolTableStack.Frames[GlobalSymbolTableStack.Length-1]
	frame.Symbols = append(frame.Symbols, s)
	frame.Length++
}

func GetVarSymbol(name string) *Symbol {
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

func GetFuncSymbol(name string) *FuncSymbol {
	for _, i := range GlobalSymbolTableStack.Funcs {
		if i.Name == name {
			return i
		}
	}
	return nil
}

func AddFunc(f *FuncSymbol) {
	GlobalSymbolTableStack.Funcs = append(GlobalSymbolTableStack.Funcs, f)
}